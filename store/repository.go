package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
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
			filter.HasPhoto == (len(parseStringInput(v.MainPhoto)) != 0) &&
			inRange(filter.CompatibilityScore, int(v.CompatibilityScore*100)) &&
			inRange(filter.Age, v.Age) &&
			inRange(filter.Height, v.HeightInCm) &&
			withinDistance(filter.DistanceInKm, v.City) {
			results = append(results, v)
		}
	}

	return results, nil
}

func inRange(r Range, value int) bool {
	return value >= r.From && value <= r.To
}

func withinDistance(r LatLngRange, c City) bool {
	p := 0.017453292519943295 // Math.PI / 180
	a := 0.5 - math.Cos((r.Lat-c.Lat)*p)/2 + math.Cos(c.Lat*p)*math.Cos(r.Lat*p)*(1-math.Cos((r.Lon-c.Lon)*p))/2
	km := 12742 * math.Asin(math.Sqrt(a)) // 2 * R; R = 6371 km

	return inRange(Range{
		From: r.From,
		To:   r.To,
	}, int(km))
}

//def distance(lat1, lon1, lat2, lon2):
//p = 0.017453292519943295
//a = 0.5 - cos((lat2 - lat1) * p)/2 + cos(lat1 * p) * cos(lat2 * p) * (1 - cos((lon2 - lon1) * p)) / 2
//return 12742 * asin(sqrt(a)) // 2 * R; R = 6371 km

func parseStringInput(input *string) string {
	if input == nil {
		return ""
	}
	return fmt.Sprintf("%v", input)
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
