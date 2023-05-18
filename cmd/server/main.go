package main

import (
	"github.com/masterkeysrd/calculation-service/internal/server"
)

func main() {
	server := server.NewServer()
	server.Start()
}
