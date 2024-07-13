package main

import (
	"github.com/krios2146/simulation-go/internal"
)

func main() {
	worldMap := internal.NewMap(10, 10)
	internal.NewSimulation(*worldMap).Start()
}
