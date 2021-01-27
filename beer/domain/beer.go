package beer_domain

// Beer ...
type Beer struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	ABV         float32 `json:"abv"`
	BreweryName string  `json:"name_breweries"`
	City        string  `json:"city"`
	Country     string  `json:"country"`
}
