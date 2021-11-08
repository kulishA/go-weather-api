package service

import (
	"github.com/kulishA/go-weather-api/domain"
	"github.com/kulishA/go-weather-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Weather interface {
	Get(user domain.User, location string) (domain.Weather, error)
	Save(user domain.User, weather domain.Weather) error
}

type Service struct {
	Authorization
	Weather
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
