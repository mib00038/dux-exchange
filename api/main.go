package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/boltdb/bolt"

	"gitlab.com/katana-labs/assessment-frontend/api/database"
	"gitlab.com/katana-labs/assessment-frontend/api/offer"
	"gitlab.com/katana-labs/assessment-frontend/api/trade"
	"gitlab.com/katana-labs/assessment-frontend/api/user"
)

func grantBudget(db *bolt.DB) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(database.BucketBalances)
		return b.Put([]byte("RGFya3dpbmcgRHVjawo="), []byte(fmt.Sprintf("%f", 500.00)))
	})
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	db := database.Init()
	defer db.Close()

	grantBudget(db)

	http.HandleFunc("/api/users/", user.NewHandler(db).Handle)
	http.HandleFunc("/api/trades/", trade.NewHandler(db).Handle)
	http.HandleFunc("/api/offers/", offer.NewHandler(ctx, db).Handle)

	log.Println("starting server on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("server stopped with error: %v", err)
	}

	cancel()

	log.Print("Goodbye")
}
