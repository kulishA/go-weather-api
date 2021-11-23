package cron

import (
	"github.com/kulishA/go-weather-api/pkg/config"
	"github.com/kulishA/go-weather-api/pkg/service"
	"github.com/sirupsen/logrus"
	"time"
)

type Cron struct {
	service *service.Service
	config  *config.Config
}

func NewCron(service *service.Service, config *config.Config) *Cron {
	return &Cron{service: service, config: config}
}

func (c *Cron) UpdateWeather() {
	updateAfter, err := time.ParseDuration(c.config.UpdateAfter)
	if err != nil {
		logrus.Error("Cron Error: can't parse .ENV var UpdateAfter")
		return
	}

	ticker := time.NewTicker(updateAfter)
	func(ticker *time.Ticker) {
		for {
			select {
			case <-ticker.C:
				logrus.Info("Start update weather")
				err := c.service.Weather.GetWeatherFromApi()
				if err != nil {
					logrus.Error("Cron Error: can't update weather")
					return
				}
			}
		}
	}(ticker)
}
