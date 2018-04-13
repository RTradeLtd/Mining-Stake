package main 

import (

	"./test_router"
	"./token_lockup"

	"encoding/binary"
	"strings"
	"bufio"
	"math/big"
	"os"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	bbolt "github.com/coreos/bbolt"

)
const key = `{"address":"069ba77207ad40b7d386f8e2979a9337a36f991c","crypto":{"cipher":"aes-128-ctr","ciphertext":"b1218c0a8d354cddcb288d021a1e76a5a8617e32b78cff0d9769b6b663851516","cipherparams":{"iv":"f1f1e9461f2e17c3ca6866173b953860"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"b2aaac5ac70f95d81b41ca1042e6ddbb1b73f456ba03c6feb8d1eda7137b571b"},"mac":"4e1eadd303f63936806808059dc1fc3be0a9706de50e5b0f824e1a4bd1310e87"},"id":"54fe5587-3f4f-45b6-b895-d905275faaf5","version":3}`
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


var m = make(map[common.Address]uint64)

func main() {

	db := bBoltSetup("stake.db")

	password := "password123"

	client, err := ethclient.Dial("/home/solidity/.ethereum/rinkeby/geth.ipc")
	if err != nil {
		log.Fatal("error connecting to network ", err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(key), password)
	if err != nil {
		log.Fatal("error connecting to network")
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("please enter address of router contract")
	scanner.Scan()
	address := scanner.Text()

	router, err := TestRouter.NewTestRouter(common.HexToAddress(address), client)
	if err != nil {
		log.Fatal("error connecting to router ", err)
	} else {
		fmt.Println(router)
	}

	fmt.Println("please enter address of lockup contract")
	scanner.Scan()
	address = scanner.Text()
	tokenLockup, err := TokenLockup.NewTokenLockup(common.HexToAddress(address), client)
	if err != nil {
		log.Fatal("error connecting to token lockup contract")
	}

	m = iterateOverBucket(db)
	var addresses []common.Address
	var rtcs      []*big.Int
	for addr, _ := range m {
		addresses = append(addresses, addr)
		hash := calculateActiveHashRate(tokenLockup, addr, db)
		exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
		dv := new(big.Int).Div(hash, exp)
		mHash := float64(dv.Int64()) / float64(1000)
		usdEarningsPerDay := float64((mHash * 1e6 / ((float64(3146) / float64(15))*1000*1e9))*((60/ float64(15))*float64(3))*(60*24)*(float64(488)))
		fmt.Println("Estimated USD earnigns a day ", usdEarningsPerDay)
		ethEarningsPerDay := float64((mHash * 1e6 / ((float64(3146) / float64(15))*1000*1e9))*((60/ float64(15))*float64(3))*(60*24))
		rtc := (usdEarningsPerDay * 0.1) / 0.125
		rtcFloat := float64(rtc)
		rtcInt := FloatToBigInt(rtcFloat)
		fmt.Println("estimated eth earnigns a day ", ethEarningsPerDay)
		fmt.Println("rtc earnings a day ", rtcInt)
		fmt.Printf("Hash rate for 0x%x\t%v\n", addr, mHash)
		rtcs = append(rtcs, rtcInt)
	}
	if len(rtcs) != len(addresses) {
		log.Fatal("not equal rtcs  || addresses")
	}
	tx, err := router.TestRouteNoRequire(auth, addresses, rtcs)
	if err != nil {
		log.Fatal("error sending transaction ", err)
	}
	fmt.Printf("transactin hash 0x%x\n", tx.Hash())
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