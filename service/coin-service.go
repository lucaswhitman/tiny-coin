package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/lucaswhitman/tiny-coin/block"
	"github.com/lucaswhitman/tiny-coin/transaction"
)

type CoinService struct {
	BlockChain            []block.Block
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

func (coinService *CoinService) Mine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address, err := vars["address"]
	if address == "" {
		RespondWithError(w, http.StatusBadRequest, "Invalid Miner Address")
		fmt.Println(err)
		return
	}

	lastBlock := coinService.BlockChain[len(coinService.BlockChain)-1]

	lastProof := lastBlock.Data.ProofOfWork
	// Find the proof of work for
	// the current block being mined
	// Note: The program will hang here until a new
	//       proof of work is found
	proof := ProofOfWork(lastProof)
	// Once we find a valid proof of work,
	// we know we can mine a block so
	// we reward the miner by adding a transaction
	t := transaction.Tranaction{"Network", address, 1}
	coinService.ThisNodesTransactions = append(coinService.ThisNodesTransactions, t)

	// Now we can gather the data needed
	// to create the new block
	newBlockData := block.Data{proof, coinService.ThisNodesTransactions}

	// Empty transaction list
	//coinService.ThisNodesTransactions := []
	// Now create the
	// new block!
	minedBlock := block.NewBlock(
		lastBlock.Index+1,
		time.Now(),
		newBlockData,
		lastBlock.Hash,
	)

	coinService.BlockChain = append(coinService.BlockChain, minedBlock)
	// Let the client know we mined a block

	fmt.Print("New block mined!\n")
	fmt.Printf("Proof of Work: %d\n", newBlockData.ProofOfWork)

	RespondWithJSON(w, http.StatusOK, minedBlock)
}

func ProofOfWork(lastProof int) int {
	// Create a variable that we will use to find
	// our next proof of work
	incrementor := lastProof + 1
	// Keep incrementing the incrementor until
	// it's equal to a number divisible by 9
	// and the proof of work of the previous
	// block in the chain

	for incrementor%9 != 0 && (lastProof == 0 || incrementor%lastProof != 0) {
		incrementor += 1
	}

	// Once that number is found,
	// we can return it as a proof
	// of our work
	return incrementor
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
