package store

import (
	"fmt"
	"testing"
)

func TestInRange(t *testing.T) {
	current := 40

	if !InRange(Range{
		From: 10,
		To:   50,
	}, current) {
		t.FailNow()
	}

	if InRange(Range{
		From: 10,
		To:   20,
	}, current) {
		t.FailNow()
	}
}

func TestParseStringInput(t *testing.T) {
	fmt.Println(ParseStringInput(nil))
	if ParseStringInput(nil) != "" {
		t.FailNow()
	}

	input := "sss"
	fmt.Println(ParseStringInput(&input))
	if ParseStringInput(&input) != input {
		t.FailNow()
	}
}

func TestWithinDistance(t *testing.T) {
	city := City{
		Name: "test",
		Lat:  51.752022,
		Lon:  -1.257677,
	}

	filter := LatLngRange{
		From: 10,
		To:   99,
		Lat:  51.509865,
		Lon:  -0.118092,
	}

	distance := GetDistance(filter, city)

	if !InRange(Range{
		From: filter.From,
		To:   filter.To,
	}, int(distance)) {
		t.FailNow()
	}
}
