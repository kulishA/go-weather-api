package api

import (
	"github.com/kulishA/go-weather-api/pkg/weaher_api/domain"
	"github.com/kulishA/go-weather-api/pkg/weaher_api/service"
)

type WeatherApi struct {
	api *service.ApiService
}

func NewWeatherApi(apiToken string) *WeatherApi {
	return &WeatherApi{
		api: service.NewApiService(apiToken),
	}
}

func (w *WeatherApi) Current(cityUrl string) (*domain.ApiWeather, error) {
	return w.api.Current(cityUrl)
}

func (w *WeatherApi) Search(cityName string) (*[]domain.ApiCity, error) {
	return w.api.Search(cityName)
}
