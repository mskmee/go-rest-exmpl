package service

import (
	"go-rest-exmpl/entities"
	"go-rest-exmpl/pkg/repository"
)

type Authorization interface {
	CreateUser(user entities.User) (string, error)
	GenerateToken(userName, password string) (string, error)
	ParseToken(token string) (string, error)
}

type TodoList interface{}

type TodoItem interface{}

type Service struct {
	TodoItem
	TodoList
	Authorization
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
	}
}
