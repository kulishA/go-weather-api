package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kulishA/go-weather-api/domain"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
}

type Weather interface {
	Get(userId int, location string) (domain.Weather, error)
	Save(userId int, weather domain.Weather) error
}

type Repository struct {
	Authorization
	Weather
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Weather:       NewWeatherPostgres(db),
	}
}
