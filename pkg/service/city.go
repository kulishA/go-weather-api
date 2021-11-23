package service

import (
	"github.com/kulishA/go-weather-api/domain"
	"github.com/kulishA/go-weather-api/pkg/repository"
	"github.com/kulishA/go-weather-api/pkg/weaher_api/api"
	apiDomain "github.com/kulishA/go-weather-api/pkg/weaher_api/domain"
)

type CityService struct {
	repo repository.City
	api  *api.WeatherApi
}

func NewCityService(repo repository.City, weatherApi *api.WeatherApi) *CityService {
	return &CityService{repo: repo, api: weatherApi}
}

func (s *CityService) Search(name string) ([]domain.City, error) {
	var result []domain.City
	cities, err := s.api.Search(name)
	if err != nil {
		return nil, err
	}

	for _, city := range *cities {
		result = append(result, s.transformCity(city))
	}

	err = s.repo.SaveAll(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *CityService) Get(userId int) ([]domain.City, error) {
	return s.repo.Get(userId)
}

func (s *CityService) Save(userId int, cityId int) (bool, error) {
	return s.repo.Save(userId, cityId)
}

func (s *CityService) transformCity(city apiDomain.ApiCity) domain.City {
	newCity := domain.City{
		CityId:  city.Id,
		Name:    city.Name,
		Region:  city.Region,
		Country: city.Country,
		Lat:     city.Lat,
		Lon:     city.Lon,
		Url:     city.Url,
	}
	return newCity
}
