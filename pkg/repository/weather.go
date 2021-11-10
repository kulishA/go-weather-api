package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kulishA/go-weather-api/domain"
)

type WeatherPostgres struct {
	db *sqlx.DB
}

func NewWeatherPostgres(db *sqlx.DB) *WeatherPostgres {
	return &WeatherPostgres{db: db}
}

func (w *WeatherPostgres) Get(userId int, location string) (domain.Weather, error) {
	var weather domain.Weather

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE location=$1 AND user_id=$2",
		weatherTable,
	)

	err := w.db.Get(&weather, query, location, userId)

	return weather, err
}

func (w *WeatherPostgres) Save(userId int, weather domain.Weather) error {
	return nil
}
