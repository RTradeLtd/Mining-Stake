package main 

import (

	"./test_router"

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

	m = iterateOverBucket(db)
	var addresses []common.Address
	var rtcs      []*big.Int
	for addr, _ := range m {
		addresses = append(addresses, addr)
		rtcs = append(rtcs, big.NewInt(500))
	}
} 