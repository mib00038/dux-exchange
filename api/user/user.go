package user

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/boltdb/bolt"

	"gitlab.com/katana-labs/assessment-frontend/api/database"
)

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Handler struct {
	db *bolt.DB
}

func NewHandler(db *bolt.DB) *Handler {
	return &Handler{db}
}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request, userID string) {
	err := h.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(database.BucketUsers)
		bts := b.Get([]byte(userID))

		if bts != nil {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(bts))
			return nil
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":"user not found"}`))

		return nil
	})
	if err != nil {
		log.Printf("could not get user from database: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"could not get user"}`))
	}
}

func (h *Handler) getBalance(w http.ResponseWriter, r *http.Request, userID string) {
	err := h.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(database.BucketBalances)
		bts := b.Get([]byte(userID))

		if bts != nil {
			balance, err := strconv.ParseFloat(string(bts), 64)
			if err != nil {
				return err
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf(`{"balance":%f}`, balance)))
			return nil
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":"balance not found"}`))

		return nil
	})
	if err != nil {
		log.Printf("could not get user from database: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"could not get user"}`))
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userID := pathParts[3]

	if len(pathParts) == 4 && r.Method == http.MethodGet {
		h.getUser(w, r, userID)
		return
	}

	action := pathParts[4]

	if action == "balance" && r.Method == http.MethodGet {
		h.getBalance(w, r, userID)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}
