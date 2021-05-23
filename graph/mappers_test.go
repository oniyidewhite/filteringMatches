package graph

import (
	"github.com/oniyidewhite/filteringMatches/graph/model"
	"github.com/oniyidewhite/filteringMatches/store"
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
		t.FailNow()
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
			Distance:           100,
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
			Distance:           100,
			MainPhoto:          nil,
			CompatibilityScore: 1,
			ContactsExchanged:  1,
			Favourite:          false,
			Religion:           "",
		},
	})

	for x := 0; x < len(test); x++ {
		a := test[x]
		b := result[x]
		if a.Favourite != b.Favourite &&
			a.Religion != b.Religion &&
			a.ContactsExchanged != b.ContactsExchanged &&
			a.Distance != b.Distance &&
			a.City != b.City &&
			a.HeightInCm != b.HeightInCm &&
			a.Age != b.Age &&
			a.CompatibilityScore != b.CompatibilityScore &&
			a.DisplayName != b.DisplayName &&
			a.City.Name != b.City.Name &&
			a.City.Lon != b.City.Lon &&
			a.City.Lat != b.City.Lat {
			t.FailNow()
		}
	}
}
