package service

import (
	"github.com/kulishA/go-weather-api/domain"
	"github.com/kulishA/go-weather-api/pkg/repository"
	"github.com/kulishA/go-weather-api/pkg/weaher_api/api"
)

const apiUrl = "https://api.weatherapi.com/v1/current.json"

type WeatherService struct {
	repo repository.Weather
	api  *api.WeatherApi
}

func NewWeatherService(repo repository.Weather) *WeatherService {
	return &WeatherService{
		repo: repo,
		api:  api.NewWeatherApi(),
	}
}

func (w *WeatherService) Get(userId int, location string) (domain.Weather, error) {
	return w.repo.Get(userId, location)
}
func (w *WeatherService) Save(userId int, weather domain.Weather) error {
	return nil
}
