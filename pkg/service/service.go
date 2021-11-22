package service

import (
	"github.com/kulishA/go-weather-api/domain"
	"github.com/kulishA/go-weather-api/pkg/repository"
	"github.com/kulishA/go-weather-api/pkg/weaher_api/api"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Weather interface {
	Get(userId int, location string) (domain.Weather, error)
	Save(userId int, weather domain.Weather) error
}

type City interface {
	Search(name string) ([]domain.City, error)
	Get(userId int) ([]domain.City, error)
	Save(userId int, cityId int) (bool, error)
}

type Service struct {
	Authorization
	Weather
	City
}

func NewService(repos *repository.Repository, api *api.WeatherApi) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Weather:       NewWeatherService(repos.Weather, api),
		City:          NewCityService(repos.City, api),
	}
}
