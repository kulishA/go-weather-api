package domain

type Weather struct {
	Id          int     `json:"-" db:"id"`
	LastUpdated string  `json:"last_updated" db:"last_updated"`
	TempC       float64 `json:"temp_c" db:"temp_c"`
	TempF       float64 `json:"temp_f" db:"temp_f"`
	IsDay       int     `json:"is_day" db:"is_day"`
	WindMph     float64 `json:"wind_mph" db:"wind_mph"`
	WindKph     float64 `json:"wind_kph" db:"wind_kph"`
	WindDegree  int     `json:"wind_degree" db:"wind_degree"`
	WindDir     string  `json:"wind_dir" db:"wind_dir"`
	PressureMb  float64 `json:"pressure_mb" db:"pressure_mb"`
	PressureIn  float64 `json:"pressure_in" db:"pressure_in"`
	PrecipMm    float64 `json:"precip_mm" db:"precip_mm"`
	PrecipIn    float64 `json:"precip_in" db:"precip_in"`
	Humidity    int     `json:"humidity" db:"humidity"`
	Cloud       int     `json:"cloud" db:"cloud"`
	FeelslikeC  float64 `json:"feelslike_c" db:"feelslike_c"`
	FeelslikeF  float64 `json:"feelslike_f" db:"feelslike_f"`
	VisKm       float64 `json:"vis_km" db:"vis_km"`
	VisMiles    float64 `json:"vis_miles" db:"vis_miles"`
	Uv          float64 `json:"uv" db:"uv"`
	GustMph     float64 `json:"gust_mph" db:"gust_mph"`
	GustKph     float64 `json:"gust_kph" db:"gust_kph"`
}

type Location struct {
	Name string `json:"name"`
}
