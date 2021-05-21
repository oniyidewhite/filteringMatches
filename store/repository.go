package store

import (
	"encoding/json"
	"io/ioutil"
)

type Repository struct {
	matches []MatchResult
}

func (r *Repository) FindMatches(filter Filter) ([]MatchResult, error) {
	return r.matches, nil
}

func NewRepository() (*Repository, error) {
	data, err := ioutil.ReadFile("matches.json")
	if err != nil {
		return nil, err
	}
	result := struct {
		Matches []MatchResult `json:"matches"`
	}{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &Repository{matches: result.Matches}, nil
}
