package zinc

import "github.com/cravenceiling/indexer/api/internal/models"

type ZincResponse struct {
	Hits struct {
		Hits []models.Hit `json:"hits"`
	} `json:"hits"`
}
