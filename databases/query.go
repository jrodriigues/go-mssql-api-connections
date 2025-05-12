package databases

import (
	"os"
)

type Query struct {
	String string
}

func NewQuery(query string) *Query {
	return &Query{query}
}

func NewQueryFromFile(filename string) (*Query, error) {
	query, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Query{string(query)}, nil
}
