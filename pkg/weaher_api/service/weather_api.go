package service

import (
	"encoding/json"
	"fmt"
	"github.com/kulishA/go-weather-api/pkg/weaher_api/domain"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	apiCurrentUrl = "https://api.weatherapi.com/v1/current.json"
	apiSearchUrl  = "https://api.weatherapi.com/v1/search.json"
	token         = "key"
	search        = "q"
)

type WeatherApiService struct {
	apiToken string
}

func NewWeatherApiService(apiToken string) *WeatherApiService {
	return &WeatherApiService{apiToken: apiToken}
}

func (w *WeatherApiService) Current(cityUrl string) (*domain.Weather, error) {
	var weather domain.Weather
	url := fmt.Sprintf("%s?%s=%s&%s=%s&aqi=no", apiCurrentUrl, token, w.apiToken, search, cityUrl)

	responseBody, err := w.request(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(responseBody, &weather)
	if err != nil {
		return nil, err
	}

	return &weather, nil
}

func (w *WeatherApiService) Search(cityName string) (*[]domain.City, error) {
	var cities []domain.City
	url := fmt.Sprintf("%s?%s=%s&%s=%s", apiSearchUrl, token, w.apiToken, search, cityName)

	responseBody, err := w.request(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(responseBody, &cities)
	if err != nil {
		return nil, err
	}

	return &cities, nil
}

func (w *WeatherApiService) request(url string) (response []byte, err error) {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	resp, err := client.Get(url)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, err
}
