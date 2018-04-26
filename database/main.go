package database

import (
	"log"

	bbolt "github.com/coreos/bbolt"
)

// BoltDB Used to hold our methods and such for teh database
type BoltDB struct {
	db *bbolt.DB
}

// Setup is used to initialize our connection to boltdb
// and create the bucket if it does not exist
func (b *BoltDB) Setup(fp string, bucket string) error {
	db, err := bbolt.Open(fp, 0600, nil)
	if err != nil {
		return err
	}
	// now that we have a valid database connection
	// we will go ahead and create the specified bucket
	// if it does not exist, otherwise we'll set the struct
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})
	if err != nil {
		log.Fatal("error establishing staker bucket in database")
	}
	b.db = db
	return nil
}
