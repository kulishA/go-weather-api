package domain

type Weather struct {
	Location   Location `json:"location"`
	TempC      float64  `json:"temp_c"`
	TempF      float64  `json:"temp_f"`
	IsDay      int      `json:"is_day"`
	WindMph    float64  `json:"wind_mph"`
	WindKph    float64  `json:"wind_kph"`
	WindDegree int      `json:"wind_degree"`
}

type Location struct {
	Name string `json:"name"`
}
