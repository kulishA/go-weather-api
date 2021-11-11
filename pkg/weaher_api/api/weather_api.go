package api

import (
	"github.com/kulishA/go-weather-api/pkg/weaher_api/domain"
	"github.com/kulishA/go-weather-api/pkg/weaher_api/service"
)

type WeatherApi struct {
	api *service.ApiService
}

func NewWeatherApi() *WeatherApi {
	//apiToken := os.Getenv("POSTGRES_PASSWORD")
	apiToken := "5a708e75b13e4239aa271034210311"
	return &WeatherApi{
		api: service.NewApiService(apiToken),
	}
}

func (w *WeatherApi) Current(cityUrl string) (*domain.Weather, error) {
	return w.api.Current(cityUrl)
}

func (w *WeatherApi) Search(cityName string) (*[]domain.City, error) {
	return w.api.Search(cityName)
}
