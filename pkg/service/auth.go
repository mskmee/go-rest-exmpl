package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"go-rest-exmpl/entities"
	"go-rest-exmpl/pkg/repository"
	"os"

	"github.com/joho/godotenv"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user entities.User) (string, error) {
	var err error
	user.Password, err = s.generatePasswordHash(user.Password)
	if err != nil {
		return "", err
	}

	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", err
	}
	salt := os.Getenv("PASSWORD_SALT")
	if salt == "" {
		return "", errors.New("PASSWORD_SALT is not set in environment variables")
	}

	hash := sha1.New()
	hash.Write([]byte(password + salt))
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
