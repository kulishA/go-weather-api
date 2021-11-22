package domain

type City struct {
	Id      int     `json:"-"`
	CityId  int     `json:"city_id" db:"city_id"`
	Name    string  `json:"name" db:"name"`
	Region  string  `json:"region" db:"region"`
	Country string  `json:"country" db:"country"`
	Lat     float64 `json:"lat" db:"lat"`
	Lon     float64 `json:"lon" db:"lon"`
	Url     string  `json:"url" db:"url"`
}
