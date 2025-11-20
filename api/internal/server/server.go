package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/cravenceiling/indexer/api/internal/config"
)

type Server struct {
	port int
}

func NewServer() *http.Server {
	port, err := strconv.Atoi(config.GetEnv("API_PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}

	NewServer := &Server{
		port: port,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.GetEnv("API_URL", ""), NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
