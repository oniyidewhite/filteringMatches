package store

import "testing"

func TestInRange(t *testing.T) {
	current := 40

	if !InRange(Range{
		From: 10,
		To:   50,
	}, current) {
		t.Failed()
	}

	if InRange(Range{
		From: 10,
		To:   20,
	}, current) {
		t.Failed()
	}
}

func TestParseStringInput(t *testing.T) {
	input := "sss"
	if ParseStringInput(&input) != input {
		t.Failed()
	}

	if ParseStringInput(nil) != "" {
		t.Failed()
	}
}

func TestWithinDistance(t *testing.T) {
	city := City{
		Name: "test",
		Lat:  51.752022,
		Lon:  -1.257677,
	}

	if !WithinDistance(LatLngRange{
		From: 10,
		To:   99,
		Lat:  51.509865,
		Lon:  -0.118092,
	}, city) {
		t.Failed()
	}
}
