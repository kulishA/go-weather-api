package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	ServerPort string
	DBPort     string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	ApiToken   string
}

func NewConfig() *Config {
	return initConfig()
}

func initConfig() *Config {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("errror loading env variables: %s", err.Error())
	}

	return &Config{
		ServerPort: os.Getenv("SERVER_PORT"),
		DBPort:     os.Getenv("POSTGRES_PORT"),
		DBHost:     os.Getenv("POSTGRES_HOST"),
		DBUser:     os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
		DBName:     os.Getenv("POSTGRES_DB"),
		DBSSLMode:  os.Getenv("POSTGRES_SSLMODE"),
		ApiToken:   os.Getenv("API_TOKEN"),
	}
}
