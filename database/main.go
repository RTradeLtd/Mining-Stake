package database

import (
	"encoding/binary"
	"fmt"
	"log"
	"math/big"

	bbolt "github.com/coreos/bbolt"
	"github.com/ethereum/go-ethereum/common"
)

// BoltDB Used to hold our methods and such for the database
// as well as our configuration information
type BoltDB struct {
	db                *bbolt.DB
	StakeIDBucketName string
	EmailBucketName   string
	OracleBucketName  string
}

var emptyString string

// Setup is used to initialize our connection to boltdb
// and create the bucket if it does not exist
func (b *BoltDB) Setup(fp string, stakeBucket string, emailBucket string, oracleBucketName string) error {
	db, err := bbolt.Open(fp, 0600, nil)
	if err != nil {
		return err
	}
	// now that we have a valid database connection
	// we will go ahead and create the specified bucket
	// if it does not exist, otherwise we'll set the struct
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(stakeBucket))
		return err
	})
	if err != nil {
		log.Fatal("error establishing staker bucket in database")
	}
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(emailBucket))
		return err
	})
	if err != nil {
		log.Fatal("error establishing email bucket")
	}
	if oracleBucketName != emptyString {
		err = db.Update(func(tx *bbolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(oracleBucketName))
			return err
		})
		if err != nil {
			log.Fatal("error establishing oracle bucket")
		}
		b.OracleBucketName = oracleBucketName
	}
	b.db = db
	b.StakeIDBucketName = stakeBucket
	b.EmailBucketName = emailBucket
	return nil
}

// UpdateOracleBucket is used to update the oracle bucket when we receive
// an oracle subscription
func (b *BoltDB) UpdateOracleBucket(address common.Address) error {
	id := big.NewInt(1)
	err := b.db.Update(func(tx *bbolt.Tx) error {
		// creates the bucket if it does not exist
		bucket, err := tx.CreateBucketIfNotExists([]byte(b.OracleBucketName))
		if err != nil {
			log.Fatal("error connecting to oracle bucket in bolt database", err)
		}
		err = bucket.Put([]byte(address.Bytes()), []byte(id.Bytes()))
		return err
	})
	if err != nil {
		return err
	}
	return nil
}

// FetchIDForOracleSubscriber is used so we can fetch the subscriber id
// if it does not exist, then we should get an empty value
func (b *BoltDB) FetchIDForOracleSubscriber(address common.Address) *big.Int {
	var response []byte
	b.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(b.OracleBucketName))
		response = bucket.Get([]byte(address.Bytes()))
		//id.SetBytes([]byte(response))
		return nil
	})
	id := new(big.Int)
	id.SetBytes(response)
	return id
}

// UpdateStakeIDBucket used to update the stakers bucket
func (b *BoltDB) UpdateStakeIDBucket(address common.Address, id *big.Int) error {
	err := b.db.Update(func(tx *bbolt.Tx) error {
		// creates the bucket if it doesnt exist
		bucket, err := tx.CreateBucketIfNotExists([]byte(b.StakeIDBucketName))
		if err != nil {
			log.Fatal("error connecting to stakers bucket in bolt database ", err)
		}
		// adds data to bucket
		err = bucket.Put([]byte(address.Bytes()), []byte(id.Bytes()))
		return err
	})
	// if err == nil, then we can assume the database wasn't updated properly
	if err != nil {
		return err
	}
	return nil
}

// RetrieveStakeIDInformationForAddress fetches the stake ID for an account
// stake id is simply the latest known stake id for an address
// we use this to iterate over the smart contract data and build
// an at run-time accurate hash rate
func (b *BoltDB) RetrieveStakeIDInformationForAddress(address common.Address) *big.Int {
	var response []byte
	b.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(b.StakeIDBucketName))
		response = bucket.Get([]byte(address.Bytes()))
		//id.SetBytes([]byte(response))
		return nil
	})
	id := new(big.Int)
	id.SetBytes(response)
	return id
}

// FetchStakeIDs used to build a list of all known stakers and their IDs
func (b *BoltDB) FetchStakeIDs() map[common.Address]uint64 {
	var m = make(map[common.Address]uint64)
	b.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(b.StakeIDBucketName))
		// Iterate over items in sorted key order.
		if err := bucket.ForEach(func(k, v []byte) error {
			address := k
			hexAddr := fmt.Sprintf("0x%x", address)
			stakeID, _ := binary.Uvarint(v)
			fmt.Printf("Staker\t0x%x\nStake ID\t%v\n", address, stakeID)
			m[common.HexToAddress(hexAddr)] = stakeID
			return nil
		}); err != nil {
			return err
		}
		return nil
	})
	return m
}

// TFetchStakeIDs used to build a list of all known stakers and their IDs for testing purposes
func (b *BoltDB) TFetchStakeIDs() (map[common.Address]uint64, error) {
	var m = make(map[common.Address]uint64)
	err := b.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(b.StakeIDBucketName))
		// Iterate over items in sorted key order.
		if err := bucket.ForEach(func(k, v []byte) error {
			address := k
			hexAddr := fmt.Sprintf("0x%x", address)
			stakeID, _ := binary.Uvarint(v)
			fmt.Printf("Staker\t0x%x\nStake ID\t%v\n", address, stakeID)
			m[common.HexToAddress(hexAddr)] = stakeID
			return nil
		}); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return m, err
	}
	return m, nil
}
