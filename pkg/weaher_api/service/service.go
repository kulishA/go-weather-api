package service

import "github.com/kulishA/go-weather-api/pkg/weaher_api/domain"

type WeatherApi interface {
	Current(cityUrl string) (*domain.Weather, error)
	Search(cityName string) (*[]domain.City, error)
}

type ApiService struct {
	WeatherApi
}

func NewApiService(apiToken string) *ApiService {
	return &ApiService{WeatherApi: NewWeatherApiService(apiToken)}
}
