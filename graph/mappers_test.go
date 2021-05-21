package graph

import (
	"github.com/oblessing/filteringMatches/graph/model"
	"github.com/oblessing/filteringMatches/store"
	"testing"
)

func TestMapInputFilter(t *testing.T) {
	result := store.Filter{
		HasPhoto:  false,
		InContact: true,
		Favourite: false,
		CompatibilityScore: store.Range{
			From: 10,
			To:   1000,
		},
		Age: store.Range{
			From: 10,
			To:   100,
		},
		Height:       store.Range{},
		DistanceInKm: store.LatLngRange{},
	}

	test := MapInputFilter(model.InputFilter{
		HasPhoto:  false,
		InContact: true,
		Favourite: false,
		CompatibilityScore: &model.Range{
			From: 10,
			To:   1000,
		},
		Age: &model.Range{
			From: 10,
			To:   100,
		},
		Height:       &model.Range{},
		DistanceInKm: &model.LatLngRange{},
	})

	if test != result {
		t.Failed()
	}
}

func TestMapMatchResult(t *testing.T) {
	result := []*model.Match{
		{
			DisplayName: "test",
			Age:         10,
			JobTitle:    "test",
			HeightInCm:  10,
			City: &model.City{
				Name: "tmp",
				Lat:  1,
				Lon:  -1,
			},
			MainPhoto:          nil,
			CompatibilityScore: 1,
			ContactsExchanged:  1,
			Favourite:          false,
			Religion:           "",
		},
	}
	test := MapMatchResult([]store.MatchResult{
		{
			DisplayName: "test",
			Age:         10,
			JobTitle:    "test",
			HeightInCm:  10,
			City: store.City{
				Name: "tmp",
				Lat:  1,
				Lon:  -1,
			},
			MainPhoto:          nil,
			CompatibilityScore: 1,
			ContactsExchanged:  1,
			Favourite:          false,
			Religion:           "",
		},
	})

	for x := 0; x < len(test); x++ {
		if test[x] != result[x] {
			t.Failed()
		}
	}
}
