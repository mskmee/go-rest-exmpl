package repository

import (
	"fmt"
	"go-rest-exmpl/entities"

	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) CreateList(userId, title, description string) (string, error) {
	var id string
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, title, description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return "", err
	}
	createUsersListQuerty := fmt.Sprintf("INSERT INTO %s (user_id, list_id) values ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuerty, userId, id)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	return id, tx.Commit()
}

func (r *TodoListPostgres) GetAllLists(userId string) ([]entities.TodoList, error) {
	var lists []entities.TodoList
	query := fmt.Sprintf("SELECT id, title, description FROM %s WHERE user_id=$1", todoListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *TodoListPostgres) GetListById(listId string) (entities.TodoList, error) {
	var list entities.TodoList
	query := fmt.Sprintf("SELECT id, title, description FROM %s WHERE id=$1", todoListsTable)
	err := r.db.Get(&list, query, listId)

	return list, err
}

func (r *TodoListPostgres) UpdateList(list entities.TodoList) error {
	query := fmt.Sprintf("UPDATE %s SET title=$1, description=$2 WHERE id=$3", todoListsTable)
	_, err := r.db.Exec(query, list.Title, list.Description, list.Id)

	return err
}

func (r *TodoListPostgres) DeleteList(listId string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", todoListsTable)
	_, err = tx.Exec(query, listId)
	if err != nil {
		tx.Rollback()
		return err
	}
	query = fmt.Sprintf("DELETE FROM %s WHERE list_id=$1", usersListsTable)
	_, err = tx.Exec(query, listId)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
