package api

import (
	"github.com/kulishA/go-weather-api/pkg/weaher_api/domain"
	"github.com/kulishA/go-weather-api/pkg/weaher_api/service"
	"reflect"
	"testing"
)

func TestNewWeatherApi(t *testing.T) {
	var tests []struct {
		name string
		want *WeatherApi
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWeatherApi(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWeatherApi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeatherApi_Current(t *testing.T) {
	type fields struct {
		api *service.ApiService
	}
	type args struct {
		cityUrl string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    domain.Weather
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WeatherApi{
				api: tt.fields.api,
			}
			got, err := w.Current(tt.args.cityUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("Current() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Current() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeatherApi_Search(t *testing.T) {
	type fields struct {
		api *service.ApiService
	}
	type args struct {
		cityName string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []domain.City
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WeatherApi{
				api: tt.fields.api,
			}
			got, err := w.Search(tt.args.cityName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
		})
	}
}
