package main

import (
	"log"

	"github.com/cravenceiling/indexer/api/internal/server"
)

func main() {
	server := server.NewServer()

	log.Println("starting server...")
	log.Printf("listening on port %s", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("cannot start server: %v", err.Error())
	}
}
