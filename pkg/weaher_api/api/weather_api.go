package api

import (
	"github.com/kulishA/go-weather-api/pkg/weaher_api/domain"
	"github.com/kulishA/go-weather-api/pkg/weaher_api/service"
	"os"
)

type WeatherApi struct {
	api *service.ApiService
}

func NewWeatherApi() *WeatherApi {

	apiToken := os.Getenv("API_TOKEN")

	return &WeatherApi{
		api: service.NewApiService(apiToken),
	}
}

func (w *WeatherApi) Current(cityUrl string) (domain.Weather, error) {
	return w.api.Current(cityUrl)
}

func (w *WeatherApi) Search(cityName string) ([]domain.City, error) {
	return w.api.Search(cityName)
}
