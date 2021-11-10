package domain

type Weather struct {
	Id         int      `json:"-" db:"id"`
	UserId     int      `json:"user_id" db:"user_id"`
	Location   Location `json:"location" db:"location"`
	TempC      float64  `json:"temp_c" db:"temp_c"`
	TempF      float64  `json:"temp_f" db:"temp_f"`
	IsDay      int      `json:"is_day" db:"is_day"`
	WindMph    float64  `json:"wind_mph" db:"wind_mph"`
	WindKph    float64  `json:"wind_kph" db:"wind_kph"`
	WindDegree int      `json:"wind_degree" db:"wind_degree"`
}

type Location struct {
	Name string `json:"name"`
}
