package repository

import (
	"go-rest-exmpl/entities"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user entities.User) (string, error)
}

type TodoList interface{}

type TodoItem interface{}

type Repository struct {
	TodoItem
	TodoList
	Authorization
}

func NewRepositories(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
