package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cravenceiling/indexer/api/internal/zinc"
)

type EmailHandler struct {
	zinc *zinc.Client
}

func NewEmailHandler() *EmailHandler {
	z := zinc.NewClient()
	if err := z.PingDB(); err != nil {
		log.Fatalf("error initializing zinc client: %v", err)
	}

	log.Println("connection to zincsearch established!")

	return &EmailHandler{
		zinc: z,
	}
}

// SearchByTerm
func (eh EmailHandler) SearchByTerm(w http.ResponseWriter, req *http.Request) {
	query, err := zinc.BuildMatchQuery(req.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}

	// Check if index exists
	index := req.URL.Query().Get("index")
	if index == "" {
		http.Error(w, "index is required", http.StatusBadRequest)
		log.Println("index is required")
		return
	}

	res, err := eh.zinc.DoZincRequest(req, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
	}
}
