package search

import "strings"

type Searchable interface {
	Terms() []string
	IsEmpty() bool
}

type searchable struct {
	query string
	terms []string
}

type SearchableRequest struct {
	Query string `json:"query" form:"query" query:"query"`
}

func NewSearchable(input SearchableRequest) Searchable {
	terms := queryToTerms(input.Query)

	return &searchable{
		query: input.Query,
		terms: terms,
	}
}

func (s *searchable) Terms() []string {
	terms := []string{}
	for _, term := range strings.Split(s.query, " ") {
		if term != "" {
			terms = append(terms, term)
		}
	}

	return terms
}

func (s *searchable) IsEmpty() bool {
	return s.query == "" || len(s.terms) == 0
}

func queryToTerms(query string) []string {
	terms := []string{}
	for _, term := range strings.Split(query, " ") {
		if term != "" {
			terms = append(terms, term)
		}
	}

	return terms
}
