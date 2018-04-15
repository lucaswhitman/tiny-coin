package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lucaswhitman/tiny-coin/transaction"
)

type CoinService struct {
	ThisNodesTransactions []transaction.Tranaction
}

func (coinService *CoinService) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var t transaction.Tranaction
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	coinService.ThisNodesTransactions = append(coinService.ThisNodesTransactions, t)

	// Because the transaction was successfully
	// submitted, we log it to our console
	fmt.Print("New transaction\n")
	fmt.Printf("FROM: %s\n", t.From)
	fmt.Printf("TO: %s\n", t.From)
	fmt.Printf("AMOUNT: %d\n", t.Amount)

	RespondWithJSON(w, http.StatusOK, t)
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
