package repository

import (
	"go-rest-exmpl/entities"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user entities.User) (string, error)
	GetUser(username, password string) (entities.User, error)
}

type TodoList interface {
	CreateList(userId, title, description string) (string, error)
	GetAllLists() ([]entities.TodoList, error)
	GetUserLists(userId string) ([]entities.TodoList, error)
	GetListById(listId string) (entities.TodoList, error)
	UpdateList(list entities.TodoList) error
	DeleteList(listId string) error
}

type TodoItem interface{}

type Repository struct {
	TodoItem
	TodoList
	Authorization
}

func NewRepositories(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoItem:      NewTodoListPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
