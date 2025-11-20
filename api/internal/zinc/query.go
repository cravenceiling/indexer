package zinc

import (
	"net/url"
	"strconv"
)

type Query struct {
	Term  string `json:"term"`
	Field string `json:"field"`
}

type matchQuery struct {
	Query      Query    `json:"query"`
	SearchType string   `json:"search_type"`
	SortFields []string `json:"sort_fields"`
	From       int      `json:"from"`
	MaxResults int      `json:"max_results"`
	Source     []string `json:"_source"`
}

func NewMatchQuery(term string) matchQuery {
	return matchQuery{
		Query: Query{
			Term:  term,
			Field: "_all",
		},
		SearchType: "match",
		SortFields: []string{"-@timestamp"},
		From:       0,
		MaxResults: 10,
		Source:     []string{},
	}
}

// buildMatchQuery
func BuildMatchQuery(params url.Values) (matchQuery, error) {
	term := params.Get("term")

	query := NewMatchQuery(term)

	page := params.Get("page")
	if page != "" {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			return matchQuery{}, err
		}
		query.From = pageInt
	}

	limit := params.Get("limit")
	if limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			return matchQuery{}, err
		}
		query.MaxResults = limitInt
	}

	return query, nil
}
