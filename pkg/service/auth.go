package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"go-rest-exmpl/entities"
	"go-rest-exmpl/pkg/repository"
	"os"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.MapClaims
	UserID string `json:"user_id"`
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

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	passwordHash, err := s.generatePasswordHash(password)
	if err != nil {
		return "", err
	}
	user, err := s.repo.GetUser(username, passwordHash)
	if err != nil {
		return "", err
	}

	if err := godotenv.Load(); err != nil {
		return "", err
	}

	secret := os.Getenv("JWT_SECRET")
	tokenLifeTime := os.Getenv("JWT_TTL")

	if secret == "" || tokenLifeTime == "" {
		return "", errors.New("JWT_SECRET or JWT_TTL is not set in environment variables")
	}

	tokenLifeTimeInt, err := strconv.Atoi(tokenLifeTime)
	if err != nil {
		return "", err
	}

	// Create the token with custom claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		MapClaims: jwt.MapClaims{
			"expiredAt": time.Now().Add(time.Duration(tokenLifeTimeInt) * time.Hour).Unix(),
			"createdAt": time.Now().Unix(),
		},
		UserID: user.Id,
	})

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *AuthService) ParseToken(stringToken string) (string, error) {
	token, err := jwt.ParseWithClaims(stringToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		if err := godotenv.Load(); err != nil {
			return nil, err
		}

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			return nil, errors.New("JWT_SECRET is not set in environment variables")
		}

		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)

	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}
