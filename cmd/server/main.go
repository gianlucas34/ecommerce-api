package main

import (
	"log"

	"github.com/gianlucas34/ecommerce-api/internal/infra/api"
)

func main() {
	server := api.NewServer()

	log.Fatal(server.Start())
}
