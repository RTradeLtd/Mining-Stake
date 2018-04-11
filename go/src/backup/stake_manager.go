package main

import (
	"log"
	"fmt"
	"strings"
	"math/big"
	"encoding/binary"
	"bufio"
	"os"
	//"strconv"
	"github.com/onrik/ethrpc"

	"github.com/howeyc/gopass"

	// bbolt will be used to store active stakers
	bbolt "github.com/coreos/bbolt"
	//prompt "github.com/c-bata/go-prompt"
	ishell "gopkg.in/abiosoft/ishell.v2"
	"github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

    //"github.com/ethereum/go-ethereum/core/types"
	//"github.com/ethereum/go-ethereum/ethstats"
	//"github.com/sendgrid/sendgrid-go"
	//"github.com/sendgrid/sendgrid-go/helpers/mail"
	
	"github.com/RTradeLtd/Mining-Stake/token_lockup"
	"./router_token"
)

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

// used to retrieve the stake ID for an account
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
	// generate new big int, and set it to start
	// compare i to end, if less than end (-1) continue, increment counter by 1
	// if 0, then we have multiple stakes to parse through
	// otherwise we just have one stake to parse
	if end.Cmp(one) == 0 {
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
		_, khSec, _, _, _, enabled, err := contract.GetStakerStruct(nil, address, zero)
		if err != nil {
			log.Fatal("error calculating active hash rate")
		}
		if enabled == true {
			return khSec
		} else {
			return zero
		}
	}
}

// WIP
func calculatePayout(mhSec float64, diffTH float64, blockTimeSec float64, blockReward float64, ethPrice float64) (float64, error) {
	//EarningsPerMonth = (UserHashMh * 1e6 / ((difficultyTH / BlockTimeSec)*1000*1e9))*((60/ BlockTimeSec)*BlockReward)*(60*24*30)*(EthPrice)
	//EarningsPerDay = (UserHashMh * 1e6 / ((difficultyTH / BlockTimeSec)*1000*1e9))*((60/ BlockTimeSec)*BlockReward)*(60*24)*(EthPrice)	
	usdEarningsPerDay := (mhSec * 1e6 / ((diffTH / blockTimeSec)*1000*1e9))*((60/ blockTimeSec)*blockReward)*(60*24)*(ethPrice)
	return usdEarningsPerDay, nil
}


// WIP (most likely will be discarded)
func buildPayoutData(contract *TokenLockup.TokenLockup, addresses []common.Address, db *bbolt.DB) {
	var m = make(map[common.Address]*big.Int)
	for i := 0; i < len(addresses); i++ {
		m[addresses[i]] = calculateActiveHashRate(contract, addresses[i], db)
	}
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

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("please enter raw contents of json key file")
	scanner.Scan()
	key := scanner.Text()

	fmt.Println("enter password to unlock key")
	pass, err := gopass.GetPasswd()
	if err != nil {
		log.Fatal("error reading password")
	}
	fmt.Println("unlocking eth account")
	auth, err := bind.NewTransactor(strings.NewReader(key), string(pass))
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


func main() {
	// create map to store stake data
	var m = make(map[common.Address]uint64)

	// setup rpc client connection
	fmt.Println("setting up rpc client")
	// setup connection to bolt database
	fmt.Println("setting up bolt database")
	db := bBoltSetup("stake.db")

	// make sure we can interact with the contract
	fmt.Println("establishing connection with contract")
	_, auth, tokenLockup := authenticateWithContract()
	_, _, tokenLockupData := authenticateWithContractForInfo()

	rpcClient := establishRpcConnection("http://127.0.0.1:8545")

	//scanner := bufio.NewScanner(os.Stdin)

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

	fmt.Println("printing block headers")
	fmt.Println(currentBlock, previousBlock)


	// big.Int type
	//diffTH_ := currentBlock.Difficulty
	//diffThInt := diffTH_.Int64()
	diffTH := float64(3088)
	totalDiffTh := currentBlock.TotalDifficulty
	currentBlockTimestamp := currentBlock.Timestamp
	previousBlockTimestamp := previousBlock.Timestamp
	blockTimeSec := currentBlockTimestamp - previousBlockTimestamp
	blockReward := float64(3)
	ethPrice := float64(400)


	fmt.Printf("Difficulty %v\nBlock Time %v\nBlock Reward %v\nEth Price %v\nTotal Diff %v\n",
		diffTH, blockTimeSec, blockReward, ethPrice, totalDiffTh)

	// used to create a new shell
	fmt.Println("establishing shell")
    shell := ishell.New()

    // display welcome info.
    shell.Println("RTrade Technology Stake Manager")

	shell.AddCmd(&ishell.Cmd{
	    Name: "single-payout",
	    Help: "construct payout data for a single staker",
	    Func: func(c *ishell.Context) {
	        // disable the '>>>' for cleaner same line input.
	        c.ShowPrompt(false)
	        defer c.ShowPrompt(true) // yes, revert after login.
	        c.Print("Address: ")
	        address := c.ReadLine()
	        hash := calculateActiveHashRate(tokenLockupData, common.HexToAddress(address), db)
	        exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	        dv := new(big.Int).Div(hash, exp)
	        mHash := float64(dv.Int64()) / float64(1000)
	        fmt.Println("mhash rate ", mHash)
	        usdEarningsPerDay, err := calculatePayout(mHash, float64(diffTH), float64(blockTimeSec), float64(blockReward), float64(ethPrice))
	        rtcPerDay := (usdEarningsPerDay * 0.1) / 0.125
	        if err != nil {
	        	log.Fatal("error parsing hash rate ", err)
	        }
	        c.Print("Mega hashes ", mHash)
	        c.Print("Usd earnings per day ", usdEarningsPerDay)
	        c.Print("RTC earnings per day ", rtcPerDay)
	    },
	})

	/*
		TO DO:
			When we check the start date, make surre a full 24 hors have past at least
	*/
	shell.AddCmd(&ishell.Cmd{
		Name: "construct-rtc-payout-data",
		Help: "build rtc payout data for active stakers",
		Func: func(c *ishell.Context) {
			var addresses []common.Address
			var rtcs      []*big.Int
			rtc := big.NewInt(45128000000000000)
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			m = iterateOverBucket(db)
			file, err := os.Create("test")
			writer := bufio.NewWriter(file)
			if err != nil {
				log.Fatal("error creating file")
			}
			for k, _ := range m {
				addresses = append(addresses, k)
				address := k
	        	hash := calculateActiveHashRate(tokenLockupData, address, db)
	        	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	        	dv := new(big.Int).Div(hash, exp)
	        	mHash := float64(dv.Int64()) / float64(1000)
	        	fmt.Println("mhash rate ", mHash)
	        	usdEarningsPerDay, err := calculatePayout(mHash, float64(diffTH), float64(blockTimeSec), float64(blockReward), float64(ethPrice))
	        	rtcPerDay := (usdEarningsPerDay * 0.1) / 0.125
	        	rtcPerDayInt := big.NewInt(int64(rtcPerDay))
	        	rtcs = append(rtcs, rtcPerDayInt)
				_, err = fmt.Fprintf(writer, "Address\t0x%x\nmHSec \t%v\nrtc per day\t%v\n", k, mHash, rtcPerDay)
				if err != nil {
					log.Fatal("error writing to file")
				}
			}
			tx, err := tokenLockup.RouteRtcRewards(auth, addresses, rtcs)
			if err != nil {
				log.Fatal("error routing token payments")
			} else {
				fmt.Println("token payments routed successfully")
				fmt.Printf("Transaction hash 0x%x\n", tx.Hash())
			}
			writer.Flush()
		},
	})


	shell.AddCmd(&ishell.Cmd{
		Name: "construct-eth-payout-data",
		Help: "build eth payout data for active stakers",
		Func: func(c *ishell.Context) {
			var addresses 	[]common.Address
			var eths 		[]*big.Int
			eth := big.NewInt(546900000000000) // eth per person
			c.ShowPrompt(false)
			// prevent prompt from showing up until we're done processing
			defer c.ShowPrompt(true)
			m := iterateOverBucket(db)
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
	        	fmt.Println("mhash rate ", mHash)
	        	usdEarningsPerDay, err := calculatePayout(mHash, float64(diffTH), float64(blockTimeSec), float64(blockReward), float64(ethPrice))
	        	fmt.Println("usd eth earnings per day ", usdEarningsPerDay)
	        	ethEarningsPerDay := float64(ethPrice) * usdEarningsPerDay
	        	ethEarningsPerDayInt := big.NewInt(int64(ethEarningsPerDay))
	        	eths = append(eths, ethEarningsPerDayInt)
	        	_, err = fmt.Println(writer, "Address\t0x%x\nmHSec\t%v\neth per day\t%v\n", k, mHash, ethEarningsPerDay)
	        	if err != nil {
	        		log.Fatal("error writing to file")
	        	}
			}
			numAddresses := big.NewInt(int64(len(addresses)))
			ethToSend := new(big.Int).Mul(numAddresses, eth)
			auth.Value = ethToSend
			tx, err := tokenLockup.RouteEthReward(auth, addresses, eth)
			if err != nil {
				log.Fatal("error routing eth payments")
			} else {
				fmt.Println("Eth payments routed successfully")
				fmt.Printf("Transaction hash 0x%x\n", tx.Hash())
			}
			writer.Flush()
		},
		})


	shell.AddCmd(&ishell.Cmd{
		Name: "iterate-over-bucket",
		Help: "iterates over the stakers bucket",
		Func: func(c *ishell.Context) {
			m = iterateOverBucket(db)
			fmt.Println(m)
			c.Print("db iter finished\n")
		},
	})

    // run shell
    shell.Run()

}
