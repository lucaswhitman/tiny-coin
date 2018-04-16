package main

import (
	"fmt"

	"github.com/lucaswhitman/tiny-coin/app"
)

var a app.App

func main() {
	fmt.Print("Server Starting...\n")
	a := app.App{}
	a.Initialize()
	fmt.Print("Server Running!!!\n")
	a.Run(8675)
}
