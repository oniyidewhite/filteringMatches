package graph

import (
	"github.com/oblessing/filteringMatches/graph/model"
	"github.com/oblessing/filteringMatches/store"
)

func MapInputFilter(filter model.InputFilter) store.Filter {
	return store.Filter{
		HasPhoto:  filter.HasPhoto,
		InContact: filter.InContact,
		Favourite: filter.Favourite,
		CompatibilityScore: store.Range{
			From: filter.CompatibilityScore.From,
			To:   filter.CompatibilityScore.To,
		},
		Age: store.Range{
			From: filter.Age.From,
			To:   filter.Age.To,
		},
		Height: store.Range{
			From: filter.Height.From,
			To:   filter.Height.To,
		},
		DistanceInKm: store.LatLngRange{
			Lat:  filter.DistanceInKm.Lat,
			Lon:  filter.DistanceInKm.Lng,
			From: filter.DistanceInKm.From,
			To:   filter.DistanceInKm.To,
		},
	}
}

func MapMatchResult(matches []store.MatchResult) []*model.Match {
	result := []*model.Match{}
	for _, c := range matches {
		result = append(result, &model.Match{
			DisplayName: c.DisplayName,
			Age:         c.Age,
			JobTitle:    c.JobTitle,
			HeightInCm:  c.HeightInCm,
			City: &model.City{
				Name: c.City.Name,
				Lat:  c.City.Lat,
				Lon:  c.City.Lon,
			},
			MainPhoto:          c.MainPhoto,
			CompatibilityScore: float64(c.CompatibilityScore),
			ContactsExchanged:  c.ContactsExchanged,
			Favourite:          c.Favourite,
			Religion:           c.Religion,
		})
	}

	return result
}
