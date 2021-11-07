package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/kulishA/go-weather-api/domains"
	"github.com/kulishA/go-weather-api/pkg/repository"
)

const salt = "awt123awt"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user domains.User) (int, error) {
	user.Password = generatePassword(user.Password)

	return s.repo.CreateUser(user)
}

func generatePassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
