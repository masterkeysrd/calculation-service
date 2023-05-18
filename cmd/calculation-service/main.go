package main

import (
	"log"

	"github.com/masterkeysrd/calculation-service/internal/server"
)

func main() {
	log.Println("Starting Calculation Service...")
	server := server.NewServer()
	server.RegisterRoutes()
	server.Start()
}
