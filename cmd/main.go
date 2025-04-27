package main

import (
	"log"

	"github.com/rahulSailesh-shah/load_balancer/internal/server"
)

func main() {
	if err := server.Run(":8080"); err != nil {
		log.Fatalf("could not start the server: %v", err)
	}
}
