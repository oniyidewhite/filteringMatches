package store

import (
	"fmt"
	"math"
)

func InRange(r Range, value int) bool {
	return value >= r.From && value <= r.To
}

func WithinDistance(r LatLngRange, c City) bool {
	p := 0.017453292519943295 // Math.PI / 180
	a := 0.5 - math.Cos((r.Lat-c.Lat)*p)/2 + math.Cos(c.Lat*p)*math.Cos(r.Lat*p)*(1-math.Cos((r.Lon-c.Lon)*p))/2
	km := 12742 * math.Asin(math.Sqrt(a)) // 2 * R; R = 6371 km

	return InRange(Range{
		From: r.From,
		To:   r.To,
	}, int(km))
}

func ParseStringInput(input *string) string {
	if input == nil {
		return ""
	}
	return fmt.Sprintf("%v", input)
}
