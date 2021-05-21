package store

import (
	"encoding/json"
)

type Repository struct {
	matches []MatchResult
}

func (r *Repository) FindMatches(filter Filter) ([]MatchResult, error) {
	results := []MatchResult{}
	for _, v := range r.matches {
		// check conditions
		if filter.Favourite == v.Favourite &&
			filter.InContact == (v.ContactsExchanged > 0) &&
			filter.HasPhoto == (len(ParseStringInput(v.MainPhoto)) != 0) &&
			InRange(filter.CompatibilityScore, int(v.CompatibilityScore*100)) &&
			InRange(filter.Age, v.Age) &&
			InRange(filter.Height, v.HeightInCm) &&
			WithinDistance(filter.DistanceInKm, v.City) {
			results = append(results, v)
		}
	}

	return results, nil
}

func NewRepository(data []byte) (*Repository, error) {
	result := struct {
		Matches []MatchResult `json:"matches"`
	}{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &Repository{matches: result.Matches}, nil
}
