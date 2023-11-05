package main

import (
	"log"

	"github.com/SergeyCherepiuk/share/client/cmd"
	"github.com/SergeyCherepiuk/share/client/pkg/clean"
)

var CleanUps = make([]func(), 0)

func main() {
	go clean.Listen()

	if err := cmd.CreateCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
