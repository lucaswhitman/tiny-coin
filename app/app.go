package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lucaswhitman/tiny-coin/transaction"

	"github.com/gorilla/mux"
	"github.com/lucaswhitman/tiny-coin/block"
	"github.com/lucaswhitman/tiny-coin/service"
)

type App struct {
	Router      *mux.Router
	coinService service.CoinService
	BlockChain  []block.Block
}

func (a *App) Initialize() {
	a.BlockChain = make([]block.Block, 0)
	a.BlockChain = append(a.BlockChain, block.CreateGenesisBlock())
	fmt.Print("App Initializing...\n")
	a.Router = mux.NewRouter()
	a.initializeRoutes()
	a.coinService = service.CoinService{a.BlockChain, make([]transaction.Tranaction, 0)}
	fmt.Print("App Initialized!!!\n")
}

func (a *App) Run(port int) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/txion", a.coinService.CreateTransaction).Methods("POST")
	a.Router.HandleFunc("/mine/{address}", a.coinService.Mine).Methods("GET")
	//a.Router.HandleFunc("/block", a.coinService.GetBlocks).Methods("PUT")

}
