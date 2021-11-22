package service

import (
	"github.com/kulishA/go-weather-api/domain"
	"github.com/kulishA/go-weather-api/pkg/repository"
	"github.com/kulishA/go-weather-api/pkg/weaher_api/api"
)

type WeatherService struct {
	repo repository.Weather
	api  *api.WeatherApi
}

func NewWeatherService(repo repository.Weather, weatherApi *api.WeatherApi) *WeatherService {
	return &WeatherService{
		repo: repo,
		api:  weatherApi,
	}
}

func (w *WeatherService) Get(userId int, location string) (domain.Weather, error) {
	return w.repo.Get(userId, location)
}
func (w *WeatherService) Save(userId int, weather domain.Weather) error {
	return nil
}
