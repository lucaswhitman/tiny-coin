package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucaswhitman/tiny-coin/service"
)

type App struct {
	Router      *mux.Router
	coinService service.CoinService
}

func (a *App) Initialize() {
	fmt.Print("App Initializing...\n")
	a.Router = mux.NewRouter()
	a.initializeRoutes()
	a.coinService = service.CoinService{}
	fmt.Print("App Initialized!!!\n")
}

func (a *App) Run(port int) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/txion", a.coinService.CreateTransaction).Methods("POST")
	//a.Router.HandleFunc("/mine", a.coinService.GetProofOfWork).Methods("GET")
	//a.Router.HandleFunc("/block", a.coinService.GetBlocks).Methods("PUT")
}
