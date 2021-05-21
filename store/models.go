package store

type Filter struct {
	HasPhoto           bool
	InContact          bool
	Favourite          bool
	CompatibilityScore Range
	Age                Range
	Height             Range
	DistanceInKm       LatLngRange
}

type Range struct {
	From int
	To   int
}

type LatLngRange struct {
	From int
	To   int
	Lat  float64
	Lon  float64
}

type MatchResult struct {
	DisplayName        string  `json:"display_name"`
	Age                int     `json:"age"`
	JobTitle           string  `json:"job_title"`
	HeightInCm         int     `json:"height_in_cm"`
	City               City    `json:"city"`
	MainPhoto          *string `json:"main_photo"`
	CompatibilityScore float32 `json:"compatibility_score"`
	ContactsExchanged  int     `json:"contacts_exchanged"`
	Favourite          bool    `json:"favourite"`
	Religion           string  `json:"religion"`
}

type City struct {
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
}
