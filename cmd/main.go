package main

import (
	weatherapi "github.com/kulishA/go-weather-api"
	"github.com/kulishA/go-weather-api/pkg/handler"
	"github.com/sirupsen/logrus"
)

func main() {

	handlers := handler.NewHandler()

	srv := new(weatherapi.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
