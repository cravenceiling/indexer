package handlers

import (
	"encoding/json"
	"fmt"
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
	query, err := zinc.BuildQuery(zinc.ZincQuery{
		Params:     req.URL.Query(),
		SearchType: zinc.MATCH_QUERY,
	})

	fmt.Println("query: ", query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
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

// GetEmails
func (eh EmailHandler) GetEmails(w http.ResponseWriter, req *http.Request) {
	query, err := zinc.BuildQuery(zinc.ZincQuery{
		Params:     req.URL.Query(),
		SearchType: zinc.MATCHALL_QUERY,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
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
