package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SergeyCherepiuk/share/server/pkg/http"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	e := http.Router()
	e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
