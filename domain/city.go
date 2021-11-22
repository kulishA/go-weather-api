package domain

type City struct {
	Id      int     `json:"-"`
	CityId  int     `json:"city_id"`
	Name    string  `json:"name"`
	Region  string  `json:"region"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Url     string  `json:"url"`
}
