package main

import (
	"log"

	"github.com/SergeyCherepiuk/share/client/cmd"
	"github.com/SergeyCherepiuk/share/client/pkg/clean"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	go clean.InterceptInterruption()
	defer clean.Clean()
	cmd.RootCmd.Execute()
}
