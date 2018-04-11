package main

import (
	"log"
	"fmt"
	"strings"
	"math/big"
	"encoding/binary"
	"bufio"
	"os"
	"time"

	"github.com/onrik/ethrpc"
	
	bbolt "github.com/coreos/bbolt"

	"github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"./token_lockup"
	"./payment_router"
)

// please change this before compiling for production
const key = `{"address":"069ba77207ad40b7d386f8e2979a9337a36f991c","crypto":{"cipher":"aes-128-ctr","ciphertext":"b1218c0a8d354cddcb288d021a1e76a5a8617e32b78cff0d9769b6b663851516","cipherparams":{"iv":"f1f1e9461f2e17c3ca6866173b953860"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"b2aaac5ac70f95d81b41ca1042e6ddbb1b73f456ba03c6feb8d1eda7137b571b"},"mac":"4e1eadd303f63936806808059dc1fc3be0a9706de50e5b0f824e1a4bd1310e87"},"id":"54fe5587-3f4f-45b6-b895-d905275faaf5","version":3}`
const password = "password123"

func bBoltSetup(dbPath string) *bbolt.DB {
	// setup the bbolt database
	db, err := bbolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatal("error opening bolt database ", err)
	}
	// create bucket if it does not exist
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("stakers"))
		return err
		})
	if err != nil {
		log.Fatal("error establishing staker bucket in database")
	}
	return db
}

// used to retrieve the latest stake ID for an account
func retrieveBucketInformationForAddress(address common.Address, db *bbolt.DB) (*big.Int) {
	// bbolt only knows []byte types, so we have to encode our response
	var response []byte
	// create a transaction request to bbolt
	// 
    db.View(func(tx *bbolt.Tx) error {
    	// initiate a bucket connection
        bucket := tx.Bucket([]byte("stakers"))
        // poll the bucket for the aprticular address in []byte format
        response = bucket.Get([]byte(address.Bytes()))
        //id.SetBytes([]byte(response))
        return nil
    })
    // establish a new big int
	i := new(big.Int)
	// copy the bytes of response, and set it to big int in `i`
	i.SetBytes(response)
	return i
}

// used to iterate over the  bucket, returning a map with the contents
func iterateOverBucket(db *bbolt.DB) map[common.Address]uint64 {
	// establish a map, key type common.Address value type uint64
	var m = make(map[common.Address]uint64)
	// create a transaction requesst to the bbolt db
	db.View(func(tx *bbolt.Tx) error {
		// establish a bucket connection
		bucket := tx.Bucket([]byte("stakers"))
	    // Iterate over every item in teh bucket, in order
	    if err := bucket.ForEach(func(k, v []byte) error {
	    	address := k
	    	// parse address into string
	    	hexAddr := fmt.Sprintf("0x%x", address)
	    	// parse the []byte value into a uint
	    	stakeId, _ := binary.Uvarint(v)
	        fmt.Printf("Staker\t0x%x\nStake ID\t%v\n", address, stakeId)
	        // update map, converting key to common.Address
	        m[common.HexToAddress(hexAddr)] = stakeId
	        return nil
	    }); err != nil {
	        return err
	    }
	    return nil
	})
	return m
}

// this is used to calculate a users currently active hash rate so we can easily factor multiple stake payments into a single payment
func calculateActiveHashRate(contract *TokenLockup.TokenLockup, address common.Address, db *bbolt.DB) *big.Int {
	// since we're dealing wtith big numbers, we have to appropriately convert
	var one = big.NewInt(1)
	var zero = big.NewInt(0)
	start := big.NewInt(0)
	end := retrieveBucketInformationForAddress(address, db)
	khSecSum := big.NewInt(0)
/*
func (x *Int) Cmp(y *Int) (r int)
    Cmp compares x and y and returns:
    -1 if x <  y
     0 if x == y
    +1 if x >  y
*/
	
	// the person only has one stake 
	if end.Cmp(zero) == 0 {
		_, khSec, _, _, _, enabled, err := contract.GetStakerStruct(nil, address, zero)
		if err != nil {
			log.Fatal("error calculating active hash rate")
		}
		if enabled == true {
			return khSec
		} else {
			return zero
		}
	} else if end.Cmp(zero) == +1 {
		// big nums can't do standard ">, ==, <" comparisons and neeed
		// tp use the built in functions into big num types
		for i := new(big.Int).Set(start); i.Cmp(end) == -1; i.Add(i, one) {
			// retrieve teh staker struct from the contract
			_, khSec, _, _, _, enabled, err := contract.GetStakerStruct(nil, address, i)
			if err != nil {
				log.Fatal("error calculcating hash rate ", err)
			}
			// if the stake is enabled, then we add it to the total.
			if enabled == true {
				khSecSum.Add(khSecSum, khSec)
			}
		}
		return khSecSum
	} else {
		// not possible, so return 0
		// we don't want to fatal panic since this would completely stop all processing
		return zero
	}
}

// used to calculate a person's estiamted USD earnings a day of mined ether
func calculateUsdEarnings(mhSec float64, diffTH float64, blockTimeSec float64, blockReward float64, ethPrice float64) (float64, error) {
	usdEarningsPerDay := (mhSec * 1e6 / ((diffTH / blockTimeSec)*1000*1e9))*((60/ blockTimeSec)*blockReward)*(60*24)*(ethPrice)
	return usdEarningsPerDay, nil
}

// used to calculate a person's estiamted USD earnings a day of mined ether
func calculateEthEarnings(mhSec float64, diffTH float64, blockTimeSec float64, blockReward float64) (float64, error) {
	ethEarningsPerDay := (mhSec * 1e6 / ((diffTH / blockTimeSec)*1000*1e9))*((60/ blockTimeSec)*blockReward)*(60*24)
	return ethEarningsPerDay, nil
}

// used to create an RPC connection with the block chain
func establishRpcConnection(rpcUrl string) *ethrpc.EthRPC {
	rpcClient := ethrpc.New(rpcUrl)
	return rpcClient
}

// authenticates with the blockchain, and the staking contract
func authenticateWithContract()  (*ethclient.Client, *bind.TransactOpts, *TokenLockup.TokenLockup) {
	fmt.Println("initiating ipc connection")
	client, err := ethclient.Dial("/home/solidity/.ethereum/rinkeby/geth.ipc")
	if err != nil {
		log.Fatal("error connecting to blockchain ", err)
	} else {
		fmt.Println("ipc connection successfully established")
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), string(password))
	if err != nil {
		log.Fatalf("error unlocking account")
	} else {
		fmt.Println("unlock successful", auth)
	}
	tokenLockup, err := TokenLockup.NewTokenLockup(common.HexToAddress("0x5Ae6C285eeB2e5a9234956cbCf9dea2C97C3A773"), client)	

	minStake, err := tokenLockup.MINSTAKE(nil)
	if err != nil {
		log.Fatal("error connecting to contract", err)
	} else {
		fmt.Println("contract connection successful, min stake ", minStake)
	}

	return client, auth, tokenLockup
}

// authenticates with the blockchain, and the payment routing contract
func authenticateWithRouter()  (*PaymentRouter.PaymentRouter) {
	fmt.Println("initiating ipc connection")
	client, err := ethclient.Dial("/home/solidity/.ethereum/rinkeby/geth.ipc")
	if err != nil {
		log.Fatal("error connecting to blockchain ", err)
	} else {
		fmt.Println("ipc connection successfully established")
	}

	paymentRouter, err := PaymentRouter.NewPaymentRouter(common.HexToAddress("0x5Ae6C285eeB2e5a9234956cbCf9dea2C97C3A773"), client)	
	if err != nil {
		log.Fatal("error establishign connecction with payment router")
	}

	return paymentRouter
}



func sendEmail(_coin string) {
	content := fmt.Sprintf("New Payment Sent for %s\n", _coin)
	from := mail.NewEmail("stake-sendgrid-api", "sgapi@rtradetechnologies.com")
	subject := "New Payments Sent"
	to := mail.NewEmail("Mining Stake", "stake@rtradetechnologies.com")

	mContent := mail.NewContent("text/html", content)
	m := mail.NewV3MailInit(from, subject, to, mContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(m)
	if err != nil {
		fmt.Println("error sending message ", err)
	} else {
		fmt.Println("message send status code ", response.StatusCode)
	}
}


func main() {

	// make sure we can use our send grid system
	if len(os.Getenv("SENDGRID_API_KEY")) <= 15 {
		log.Fatal("no valid send grid api key detected in environment variable, please enter your api key then re-run")
	}

	currDate := time.Now()
	weekday := currDate.Weekday()
	weekdayString := weekday.String()

	// conecting to bbolt data base
	db := bBoltSetup("stake.db")

	// establishing connection with stake contract
	_, auth, tokenLockup := authenticateWithContract()
	paymentRouter := authenticateWithRouter()

	// connecting to rpc backend
	rpcClient := establishRpcConnection("http://127.0.0.1:8545")

	currentBlockNum, err := rpcClient.EthBlockNumber()
	if err != nil {
		log.Fatal("error reading block num ", err)
	}

	previousBlockNum := currentBlockNum - 1

	// now that we have the latest and previous block, we can go about
	// parsing the data

	currentBlock, err := rpcClient.EthGetBlockByNumber(currentBlockNum, false)
	if err != nil {
		log.Fatal("error retrieving current  block headers ",err)
	}

	previousBlock, err := rpcClient.EthGetBlockByNumber(previousBlockNum, false)
	if err != nil {
		log.Fatal("error retrieving previous block headers ", err)
	}

	diffTH := float64(3088)
	currentBlockTimestamp := currentBlock.Timestamp
	previousBlockTimestamp := previousBlock.Timestamp
	blockTimeSec := currentBlockTimestamp - previousBlockTimestamp
	blockReward := float64(3)
	ethPrice := float64(419)

	if weekdayString == "Friday" {
		var m = make(map[common.Address]uint64)
		// today is friday so we will also send out eth
		// lets start with the eth payment
		var addresses 	[]common.Address
		var eths 		[]*big.Int
		eth := big.NewInt(5000)
		m = iterateOverBucket(db)
		file, err := os.Create("eth.txt")
		writer := bufio.NewWriter(file)
		if err != nil {
			log.Fatal("error creating file")
		}
		for k, _ := range m {
			addresses = append(addresses, k)
			address := k
			hash := calculateActiveHashRate(tokenLockup, address, db)
			exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	        dv := new(big.Int).Div(hash, exp)
	        mHash := float64(dv.Int64()) / float64(1000)
        	ethEarningsPerDay, err := calculateEthEarnings(mHash, float64(diffTH), float64(blockTimeSec), float64(blockReward))
        	ethEarningsPerDayInt := int64(ethEarningsPerDay)
        	ethEarningsPerDayBigInt := big.NewInt(ethEarningsPerDayInt)
        	if err != nil {
        		fmt.Println("error reading eth earnings ", err)
        	}
	        eths = append(eths, ethEarningsPerDayBigInt)
	        _, err = fmt.Println(writer, "Address\t0x%x\nmHSec\t%v\neth per day\t%v\n", k, mHash, ethEarningsPerDay)
	        if err != nil {
	        	log.Fatal("error writing to file")
	        }
			numAddresses := big.NewInt(int64(len(addresses)))
			totalEthToSend := new(big.Int).Mul(numAddresses, eth)
			auth.Value = totalEthToSend
			tx, err := paymentRouter.RouteEthReward(auth, addresses, eths)
			if err != nil {
				log.Fatal("error routing eth payments")
			} else {
				fmt.Println("Eth payments routed successfully")
				fmt.Printf("Transaction hash 0x%x\n", tx.Hash())
				sendEmail("eth")
			}
			writer.Flush()
		}
	}

	var m = make(map[common.Address]uint64)
	var addresses 	[]common.Address
	var rtcs		[]*big.Int
	m = iterateOverBucket(db)
	file, err := os.Create("rtc.txt")
	if err != nil {
		log.Fatal("error creating file")
	}
	writer := bufio.NewWriter(file)
	for k, _ := range m {
		addresses = append(addresses, k)
		address := k
		hash := calculateActiveHashRate(tokenLockup, address, db)
		exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	    dv := new(big.Int).Div(hash, exp)
	    mHash := float64(dv.Int64()) / float64(1000)
		usdEarningsPerDay, err := calculateUsdEarnings(mHash, float64(diffTH), float64(blockTimeSec), float64(blockReward), float64(ethPrice))
	    rtcPerDay := (usdEarningsPerDay * 0.1) / 0.125
	    rtcPerDayInt := big.NewInt(int64(rtcPerDay))
	    rtcs = append(rtcs, rtcPerDayInt)
		_, err = fmt.Fprintf(writer, "Address\t0x%x\nmHSec \t%v\nrtc per day\t%v\n", k, mHash, rtcPerDay)
		if err != nil {
			log.Fatal("error writing to file")
		}
	}
	tx, err := paymentRouter.RouteRtcRewards(auth, addresses, rtcs)
	if err != nil {
		log.Fatal("error routing token payments")
	} else {
		fmt.Println("token payments routed successfully")
		fmt.Printf("Transaction hash 0x%x\n", tx.Hash())
		sendEmail("rtc")
	}
	writer.Flush()

}