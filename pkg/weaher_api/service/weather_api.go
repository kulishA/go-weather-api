package service

import "github.com/kulishA/go-weather-api/pkg/weaher_api/domain"

type WeatherApiService struct {
	apiToken string
}

func NewWeatherApiService(apiToken string) *WeatherApiService {
	return &WeatherApiService{apiToken: apiToken}
}

func (w *WeatherApiService) Current(cityUrl string) (domain.Weather, error) {
	return domain.Weather{}, nil
}

func (w *WeatherApiService) Search(cityName string) ([]domain.City, error) {
	return []domain.City{}, nil
}
