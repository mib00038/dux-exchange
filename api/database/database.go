package database

import (
	"log"

	"github.com/boltdb/bolt"
)

var (
	BucketUsers    = []byte("users")
	BucketBalances = []byte("balances")
	BucketOffers   = []byte("offers")
	Buckets        = [][]byte{
		BucketUsers,
		BucketBalances,
		BucketOffers,
	}
)

func Init() *bolt.DB {
	db, err := bolt.Open("dux.db", 0600, nil)
	if err != nil {
		log.Fatalf("could not open database: %v", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		for _, b := range Buckets {
			_, err := tx.CreateBucketIfNotExists(b)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		log.Fatalf("could not create buckets in database: %v", err)
	}

	return db
}
