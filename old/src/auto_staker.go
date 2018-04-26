package main 

import (

	"./token_lockup"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"encoding/binary"
	"strconv"
	"strings"
	"bufio"
	"math/big"
	"os"
	"fmt"
	"log"
	"time"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/howeyc/gopass"
	"github.com/onrik/ethrpc"
	bbolt "github.com/coreos/bbolt"

)

const key = `{"address":"069ba77207ad40b7d386f8e2979a9337a36f991c","crypto":{"cipher":"aes-128-ctr","ciphertext":"b1218c0a8d354cddcb288d021a1e76a5a8617e32b78cff0d9769b6b663851516","cipherparams":{"iv":"f1f1e9461f2e17c3ca6866173b953860"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"b2aaac5ac70f95d81b41ca1042e6ddbb1b73f456ba03c6feb8d1eda7137b571b"},"mac":"4e1eadd303f63936806808059dc1fc3be0a9706de50e5b0f824e1a4bd1310e87"},"id":"54fe5587-3f4f-45b6-b895-d905275faaf5","version":3}`


type Response struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Symbol string `json:"symbol"`
	Rank string `json:"rank"`
	PriceUsd string `json:"price_usd"`
	PriceBtc string `json:"price_btc"`
	TwentyFourHrVolume string `json:"24h_volume_usd"`
	MarketCapUsd string `json:"market_cap_usd"`
	AvailableSupply string `json:"available_supply"`
	TotalSupply string `json:"total_supply"`
	MaxSupply string `json:"null"`
	PercentChange1h string `json:"percent_change_1h"`
	PercentChange24h string `json:"percent_change_24h"`
	PercentChange7d string `json:"percent_change_7d"`
	LastUpdate string `json:"last_updated"`
}

var m = make(map[common.Address]uint64)

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

// used to update the database
func updateBboltDb(address common.Address, id *big.Int, db *bbolt.DB) {
	err := db.Update(func(tx *bbolt.Tx) error {
		// creates the bucket if it doesnt exist
		bucket, err := tx.CreateBucketIfNotExists([]byte("stakers"))
		if err != nil {
			log.Fatal("error connecting to stakers bucket in bolt database ", err)
		}
		// adds data to bucket
		err = bucket.Put([]byte(address.Bytes()), []byte(id.Bytes()))
		return err
	})
	// if err == nil, then we can assume the database wasn't updated properly
	if err != nil {
		log.Fatal("error updating bolt stakers bucket ", err)
	}
}

// used to retrieve the stake ID for an account
func retrieveBucketInformationForAddress(address common.Address, db *bbolt.DB) (*big.Int) {
	var response []byte
    db.View(func(tx *bbolt.Tx) error {
        bucket := tx.Bucket([]byte("stakers"))
        response = bucket.Get([]byte(address.Bytes()))
        //id.SetBytes([]byte(response))
        return nil
    })
	i := new(big.Int)
	i.SetBytes(response)
	return i
}



// this is used to calculate a users currently active hash rate so we can easily factor multiple stake payments into a single payment
func calculateActiveHashRate(contract *TokenLockup.TokenLockup, address common.Address, db *bbolt.DB) *big.Int {
	var one = big.NewInt(1)
	var zero = big.NewInt(0)
	start := big.NewInt(0)
	end := retrieveBucketInformationForAddress(address, db)
	khSecSum := big.NewInt(0)
	// generate new big int, and set it to start
	// compare i to end, if less than end (-1) continue, increment counter by 1
	if end.Cmp(one) == 0 {
		for i := new(big.Int).Set(start); i.Cmp(end) == -1; i.Add(i, one) {
			_, khSec, _, _, _, enabled, err := contract.GetStakerStruct(nil, address, i)
			if err != nil {
				log.Fatal("error calculcating hash rate ", err)
			}
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


// used to iterate over the  bucket, returning a map with the contents
func iterateOverBucket(db *bbolt.DB) map[common.Address]uint64 {
	var m = make(map[common.Address]uint64)
	db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("stakers"))
	    // Iterate over items in sorted key order.
	    if err := bucket.ForEach(func(k, v []byte) error {
	    	address := k
	    	hexAddr := fmt.Sprintf("0x%x", address)
	    	stakeId, _ := binary.Uvarint(v)
	        fmt.Printf("Staker\t0x%x\nStake ID\t%v\n", address, stakeId)
	        m[common.HexToAddress(hexAddr)] = stakeId
	        return nil
	    }); err != nil {
	        return err
	    }
	    return nil
	})
	return m
}

func parseCmcApi() float64 {

	response, err := http.Get("https://api.coinmarketcap.com/v1/ticker/ethereum/")
	if err != nil {
		log.Fatal("error reading website ", err)
	}


	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("error reading response ", err)
	}
	var decode []Response
	err = json.Unmarshal(body, &decode)
	if err != nil {
		log.Fatal("error unmarshling json ", err)
	}

	f, _ := strconv.ParseFloat(decode[0].PriceUsd, 64)

	return f
}

func calculateUsdPayout(mhSec float64, diffTH float64, blockTimeSec float64, blockReward float64, ethPrice float64) (float64, error) {
	//EarningsPerMonth = (UserHashMh * 1e6 / ((difficultyTH / BlockTimeSec)*1000*1e9))*((60/ BlockTimeSec)*BlockReward)*(60*24*30)*(EthPrice)
	//EarningsPerDay = (UserHashMh * 1e6 / ((difficultyTH / BlockTimeSec)*1000*1e9))*((60/ BlockTimeSec)*BlockReward)*(60*24)*(EthPrice)	
	usdEarningsPerDay := float64((mhSec * 1e6 / ((diffTH / blockTimeSec)*1000*1e9))*((60/ blockTimeSec)*blockReward)*(60*24)*(ethPrice))
	return usdEarningsPerDay, nil
}

func calculateEthPayout(mhSec float64, diffTH float64, blockTimeSec float64, blockReward float64) (float64, error) {
	//EarningsPerMonth = (UserHashMh * 1e6 / ((difficultyTH / BlockTimeSec)*1000*1e9))*((60/ BlockTimeSec)*BlockReward)*(60*24*30)*(EthPrice)
	//EarningsPerDay = (UserHashMh * 1e6 / ((difficultyTH / BlockTimeSec)*1000*1e9))*((60/ BlockTimeSec)*BlockReward)*(60*24)*(EthPrice)	
	usdEarningsPerDay := float64((mhSec * 1e6 / ((diffTH / blockTimeSec)*1000*1e9))*((60/ blockTimeSec)*blockReward)*(60*24))
	return usdEarningsPerDay, nil
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

// used to create an RPC connection with the block chain
func establishRpcConnection(rpcUrl string) *ethrpc.EthRPC {
	rpcClient := ethrpc.New(rpcUrl)
	return rpcClient
}

func main() {

	// setup rpc client connection
	fmt.Println("setting up rpc client")
	// setup connection to bolt database
	fmt.Println("setting up bolt database")
	db := bBoltSetup("stake.db")

	// make sure we can interact with the contract
	fmt.Println("establishing connection with contract")
	_, auth, tokenLockup := authenticateWithContract()

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

	fmt.Println("printing block headers")
	fmt.Println(currentBlock, previousBlock)
	// big.Int type
	//diffTH_ := currentBlock.Difficulty
	//diffThInt := diffTH_.Int64()
	diffTH := float64(3200)
	currentBlockTimestamp := currentBlock.Timestamp
	previousBlockTimestamp := previousBlock.Timestamp
	blockTimeSec := currentBlockTimestamp - previousBlockTimestamp
	blockReward := float64(3)
	ethPrice := parseCmcApi()
	

	currDate := time.Now()
	weekday := currDate.Weekday()
	// calculate eth payments
	if weekday.String() == "Saturday" {
		var m = make(map[common.Address]uint64)
		m = iterateOverBucket(db)
		for addr, _ := range m {
			var address []common.Address
			address = append(address, addr)
			hash := calculateActiveHashRate(tokenLockup, addr, db)
			exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
			dv := new(big.Int).Div(hash, exp)
			mHash := float64(dv.Int64()) / float64(1000)
			ethEarnings, _ := calculateEthPayout(mHash, float64(diffTH), float64(blockTimeSec), float64(blockReward))
			ethEarningsBig := FloatToBigInt(ethEarnings)
			weekEarnings := new(big.Int).Mul(ethEarningsBig, big.NewInt(7))
			auth.Value = weekEarnings
			tx, err := tokenLockup.RouteEthReward(auth, address, weekEarnings)
			if err != nil {
				log.Fatal("error sending token ", err)
			} 
			fmt.Printf("TX Hash 0x%x\n", tx.Hash())
		}
	}
	log.Fatal()
	auth.Value = big.NewInt(0)

	var m = make(map[common.Address]uint64)
	m = iterateOverBucket(db)
	for k, _ := range m {
		var address []common.Address
		address = append(address, k)
		hash := calculateActiveHashRate(tokenLockup, k, db)
		// since we're dealing with big numbers, we can simply just divide by 10^18, we need to do that by utilizing big int variables
		exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
		dv := new(big.Int).Div(hash, exp)
		mHash := float64(dv.Int64()) / float64(1000)
		usdEarnings, _ := calculateUsdPayout(mHash, float64(diffTH), float64(blockTimeSec), float64(blockReward), float64(ethPrice))
		percentUsd := new(big.Float).Mul(big.NewFloat(usdEarnings), big.NewFloat(0.1))
		percentUsdFloat, _ := percentUsd.Float64()
		fmt.Println("USD Float ", percentUsdFloat)
		rtcFloat := percentUsdFloat / 0.125
		rtc := FloatToBigInt(rtcFloat)
		fmt.Println("rtc ", rtc)
		tx, err := tokenLockup.RouteRtcRewards(auth, address, rtc)
		if err != nil {
			log.Fatal("error routing token payments ", err)
		} else {
			fmt.Println("token payments routed successfully")
			fmt.Printf("Transaction hash 0x%x\n", tx.Hash())
		}
	}

} 

func FloatToBigInt(val float64) *big.Int {
    bigval := new(big.Float)
    bigval.SetFloat64(val)
    // Set precision if required.
    // bigval.SetPrec(64)

    coin := new(big.Float)
    coin.SetInt(big.NewInt(1000000000000000000))

    bigval.Mul(bigval, coin)

    result := new(big.Int)
    bigval.Int(result) // store converted number in result

    return result
}