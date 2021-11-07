package service

import (
	"github.com/kulishA/go-weather-api/domains"
	"github.com/kulishA/go-weather-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user domains.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
