package offer

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/boltdb/bolt"

	"gitlab.com/katana-labs/assessment-frontend/api/database"
	"gitlab.com/katana-labs/assessment-frontend/api/duck"
)

type Offer struct {
	ID        string    `json:"id"`
	UnitType  duck.Duck `json:"unitType"`
	UnitPrice float64   `json:"unitPrice"`
	Volume    int       `json:"volume"`
	CreatedOn time.Time `json:"createdOn"`
}

func Rand() Offer {
	return Offer{
		ID:        strconv.Itoa(rand.Intn(999999999999) + 100000000000),
		UnitType:  duck.Rand(),
		UnitPrice: float64(rand.Intn(2000)+20) / 100,
		Volume:    rand.Intn(300) + 1,
		CreatedOn: time.Now(),
	}
}

type Handler struct {
	ctx    context.Context
	db     *bolt.DB
	offers chan Offer
}

func NewHandler(ctx context.Context, db *bolt.DB) *Handler {
	offers := make(chan Offer)

	handler := &Handler{ctx, db, offers}

	go handler.startSteam()
	go handler.startVacuuming()

	return handler
}

func (h *Handler) startSteam() {
	for {
		select {
		case <-h.ctx.Done():
			return
		case <-time.After(time.Duration(rand.Intn(5)) * time.Second):
			offer := Rand()

			bts, err := json.Marshal(offer)
			if err != nil {
				log.Printf("could not marshal offer to JSON: %v", err)
				continue
			}

			err = h.db.Update(func(tx *bolt.Tx) error {
				b := tx.Bucket(database.BucketOffers)

				err := b.Put([]byte(offer.ID), bts)
				if err != nil {
					return err
				}

				h.offers <- offer

				return nil
			})
			if err != nil {
				log.Printf("could not store offer in database: %v", err)
			}
		}
	}
}

func (h *Handler) startVacuuming() {
	for {
		select {
		case <-h.ctx.Done():
			return
		case <-time.After(10 * time.Second):
			tenSecondsAgo := time.Now().Add(-10 * time.Second)

			err := h.db.Update(func(tx *bolt.Tx) error {
				b := tx.Bucket(database.BucketOffers)

				return b.ForEach(func(k, v []byte) error {
					var offer Offer
					err := json.Unmarshal(v, &offer)
					if err != nil {
						return err
					}

					if offer.CreatedOn.Before(tenSecondsAgo) {
						return b.Delete([]byte(offer.ID))
					}

					return nil
				})
			})
			if err != nil {
				log.Printf("could not vacuum offers: %v", err)
			}
		}
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Content-Type", "application/json")

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-h.ctx.Done():
				done <- true
				return
			case <-time.After(1 * time.Minute):
				w.Write([]byte(`{"status":"timeout"}`))
				done <- true
				return
			case offer := <-h.offers:
				bts, err := json.Marshal(offer)
				if err != nil {
					log.Printf("could not marshal offer to JSON: %v", err)
					done <- true
					return
				}

				w.Write([]byte(bts))

				done <- true
				return
			}
		}
	}()

	<-done
}
