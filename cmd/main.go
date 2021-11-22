package main

import (
	"github.com/kulishA/go-weather-api/pkg/config"
	"github.com/kulishA/go-weather-api/pkg/handler"
	"github.com/kulishA/go-weather-api/pkg/repository"
	"github.com/kulishA/go-weather-api/pkg/server"
	"github.com/kulishA/go-weather-api/pkg/service"
	"github.com/kulishA/go-weather-api/pkg/weaher_api/api"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	configs := config.NewConfig()

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     configs.DBHost,
		Port:     configs.DBPort,
		Username: configs.DBUser,
		Password: configs.DBPassword,
		DBName:   configs.DBName,
		SSLMode:  configs.DBSSLMode,
	})

	if err != nil {
		logrus.Fatalf("failded to initialize DB: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	weatherApi := api.NewWeatherApi(configs.ApiToken)
	services := service.NewService(repos, weatherApi)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(configs.ServerPort, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
