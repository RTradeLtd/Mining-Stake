package main

import (
	"fmt"
	"log"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
	bbolt "github.com/coreos/bbolt"
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

func updateBboltDb(address common.Address, id *big.Int, db *bbolt.DB) {
	err := db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("stakers"))
		if err != nil {
			log.Fatal("error connecting to stakers bucket in bolt database ", err)
		}
		err = bucket.Put([]byte(address.Bytes()), []byte(id.Bytes()))
		return err
	})
	if err != nil {
		log.Fatal("error updating bolt stakers bucket ", err)
	}
}

func retrieveBoltInformationForAddress(address common.Address, db *bbolt.DB) (id big.Int) {
	db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("stakers"))
		response := bucket.Get([]byte(address.Bytes()))
		id.SetBytes([]byte(response))
		return nil
	})
	return id
}

func main() {
	db := bBoltSetup("stakers.db")
	fmt.Println("database connected to successfully")
	updateBboltDb(common.HexToAddress("0x"), big.NewInt(100), db)
	fmt.Println("database written to successfully")
	id := retrieveBolInformationForAddress(common.HexToAddress("0x"), db)
	fmt.Println("id ", id)
}
