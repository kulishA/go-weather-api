package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kulishA/go-weather-api/domain"
)

type WeatherPostgres struct {
	db *sqlx.DB
}

func NewWeatherPostgres(db *sqlx.DB) *WeatherPostgres {
	return &WeatherPostgres{db: db}
}

func (w *WeatherPostgres) Get(user domain.User, location string) (domain.Weather, error) {
	return domain.Weather{}, nil
}

func (w *WeatherPostgres) Save(user domain.User, weather domain.Weather) error {
	return nil
}
