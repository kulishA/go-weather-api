package service

import (
	"github.com/kulishA/go-weather-api/domain"
	"github.com/kulishA/go-weather-api/pkg/repository"
	"github.com/kulishA/go-weather-api/pkg/weaher_api/api"
)

type WeatherService struct {
	rw  repository.Weather
	rc  repository.City
	api *api.WeatherApi
}

func NewWeatherService(rw repository.Weather, rc repository.City, weatherApi *api.WeatherApi) *WeatherService {
	return &WeatherService{
		rw:  rw,
		rc:  rc,
		api: weatherApi,
	}
}

func (w *WeatherService) Get(userId int, cityId int) (domain.Weather, error) {
	return w.rw.Get(userId, cityId)
}

func (w *WeatherService) GetAll(userId int) ([]domain.Weather, error) {
	return w.rw.GetAll(userId)
}

func (w *WeatherService) GetWeatherFromApi() error {
	cities, err := w.rc.GetAllSaved()
	if err != nil {
		return err
	}

	for _, city := range cities {
		city := city
		go func() {
			weather, _ := w.api.Current(city.Url)
			err := w.rw.SaveOrUpdate(city.CityId, *weather)
			if err != nil {
				return
			}
		}()
	}

	return nil
}
