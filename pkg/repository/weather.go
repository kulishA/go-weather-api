package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kulishA/go-weather-api/domain"
	apiDomain "github.com/kulishA/go-weather-api/pkg/weaher_api/domain"
)

type WeatherPostgres struct {
	db *sqlx.DB
}

func NewWeatherPostgres(db *sqlx.DB) *WeatherPostgres {
	return &WeatherPostgres{db: db}
}

func (w *WeatherPostgres) Get(userId int, cityId int) (domain.Weather, error) {
	var weather domain.Weather

	query := fmt.Sprintf("SELECT DISTINCT last_updated, temp_c, temp_f, is_day, wind_mph, wind_kph, " +
		"wind_degree, wind_dir, pressure_mb, pressure_in, precip_mm, precip_in, humidity, cloud, feelslike_c, " +
		"feelslike_f, vis_km, vis_miles, uv, gust_mph, gust_kph FROM %s AS w " +
		"INNER JOIN %s AS uc ON uc.user_id=$1 AND uc.city_id = $2 AND uc.city_id = w.city_id", weatherTable, usersCitiesTable)
	err := w.db.Get(&weather, query, userId, cityId)

	return weather, err
}

func (w *WeatherPostgres) GetAll(userId int) ([]domain.Weather, error) {
	var weather []domain.Weather

	query := fmt.Sprintf("SELECT DISTINCT last_updated, temp_c, temp_f, is_day, wind_mph, wind_kph, " +
		"wind_degree, wind_dir, pressure_mb, pressure_in, precip_mm, precip_in, humidity, cloud, feelslike_c, " +
		"feelslike_f, vis_km, vis_miles, uv, gust_mph, gust_kph FROM %s AS w " +
		"INNER JOIN %s AS uc ON uc.user_id=$1 AND uc.city_id = w.city_id", weatherTable, usersCitiesTable)
	err := w.db.Select(&weather, query, userId)

	return weather, err
}

func (w *WeatherPostgres) SaveOrUpdate(cityId int, weather apiDomain.ApiWeather) error {
	query := fmt.Sprintf("INSERT INTO %s (city_id, last_updated, temp_c, temp_f, is_day, wind_mph, wind_kph, "+
		"wind_degree, wind_dir, pressure_mb, pressure_in, precip_mm, precip_in, humidity, cloud, feelslike_c,"+
		"feelslike_f, vis_km, vis_miles, uv, gust_mph, gust_kph) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12,"+
		"$13, $14, $15, $16, $17, $18, $19, $20, $21, $22)", weatherTable)

	query += " ON CONFLICT (city_id) DO UPDATE SET last_updated=$2, temp_c=$3, temp_f=$4, is_day=$5, wind_mph=$6," +
		"wind_kph=$7, wind_degree=$8, wind_dir=$9, pressure_mb=$10, pressure_in=$11, precip_mm=$12, precip_in=$13, humidity=$14," +
		"cloud=$15, feelslike_c=$16, feelslike_f=$17, vis_km=$18, vis_miles=$19, uv=$20, gust_mph=$21, gust_kph=$22"

	_, err := w.db.Exec(query, cityId, weather.LastUpdated, weather.TempC, weather.TempF,
		weather.IsDay, weather.WindMph, weather.WindKph, weather.WindDegree, weather.WindDir,
		weather.PressureMb, weather.PressureIn, weather.PrecipMm, weather.PrecipIn, weather.Humidity,
		weather.Cloud, weather.FeelslikeC, weather.FeelslikeF, weather.VisKm, weather.VisMiles, weather.Uv,
		weather.GustMph, weather.GustKph)

	if err != nil {
		return err
	}
	return nil
}
