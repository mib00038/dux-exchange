package trade

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"gitlab.com/katana-labs/assessment-frontend/api/database"
	"gitlab.com/katana-labs/assessment-frontend/api/offer"
)

var (
	errVolumeTooSmall    = errors.New("volume too small")
	errVolumeTooLarge    = errors.New("volume too large")
	errOfferExpired      = errors.New("offer expired")
	errInsufficientFunds = errors.New("insufficient funds")
)

type Trade struct {
	UserID  string `json:"userId"`
	OfferID string `json:"offerId"`
	Volume  int    `json:"volume"`
}

type Handler struct {
	db *bolt.DB
}

func NewHandler(db *bolt.DB) *Handler {
	return &Handler{db}
}

func validateTrade(timeLimit time.Time, o offer.Offer, t Trade, balance float64) error {
	if t.Volume <= 0 {
		return errVolumeTooSmall
	}
	if t.Volume > o.Volume {
		return errVolumeTooLarge
	}
	if o.CreatedOn.After(timeLimit) {
		return errOfferExpired
	}

	totalPrice := float64(t.Volume) * o.UnitPrice

	if totalPrice > balance {
		return errInsufficientFunds
	}

	return nil
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	tenSecondsAgo := time.Now().Add(-10 * time.Second)

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var trade Trade
	err := json.NewDecoder(r.Body).Decode(&trade)
	if err != nil {
		log.Printf("could not parse body as trade: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"body not a valid trade"}`))
		return
	}

	err = h.db.Update(func(tx *bolt.Tx) error {
		bb := tx.Bucket(database.BucketBalances)
		ob := tx.Bucket(database.BucketOffers)

		balanceBts := bb.Get([]byte(trade.UserID))
		if balanceBts == nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"balance not found"}`))
			return nil
		}

		balance, err := strconv.ParseFloat(string(balanceBts), 64)
		if err != nil {
			return err
		}

		offerBts := ob.Get([]byte(trade.OfferID))
		if offerBts == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error":"offer not found, expired or traded"}`))
			return nil
		}

		var offer offer.Offer
		err = json.Unmarshal(offerBts, &offer)
		if err != nil {
			return err
		}

		err = validateTrade(tenSecondsAgo, offer, trade, balance)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
			return nil
		}

		totalPrice := float64(trade.Volume) * offer.UnitPrice
		newBalance := balance - totalPrice

		err = bb.Put([]byte(trade.UserID), []byte(fmt.Sprintf("%f", newBalance)))
		if err != nil {
			return err
		}

		w.WriteHeader(http.StatusAccepted)

		return nil
	})
	if err != nil {
		log.Printf("could not execute trade: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"could not execute trade"}`))
	}
}
