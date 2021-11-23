package service

import "github.com/kulishA/go-weather-api/pkg/weaher_api/domain"

type WeatherApi interface {
	Current(cityUrl string) (*domain.ApiWeather, error)
	Search(cityName string) (*[]domain.ApiCity, error)
}

type ApiService struct {
	WeatherApi
}

func NewApiService(apiToken string) *ApiService {
	return &ApiService{WeatherApi: NewWeatherApiService(apiToken)}
}
