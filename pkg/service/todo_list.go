package service

import (
	"go-rest-exmpl/entities"
	"go-rest-exmpl/pkg/repository"
)

type TodoService struct {
	repository repository.TodoList
}

func NewTodoListService(repository repository.TodoList) *TodoService {
	return &TodoService{repository: repository}
}

func (s *TodoService) CreateList(userId string, list entities.TodoList) (string, error) {
	return s.repository.CreateList(userId, list.Title, list.Description)
}

// func (s *TodoService) GetLists(userId string) ([]entities.TodoList, error) {
// 	return s.repository.GetAllLists(userId)
// }

// func (s *TodoService) GetListById(listId string) (entities.TodoList, error) {
// 	return s.repository.GetListById(listId)
// }

// func (s *TodoService) UpdateList(list entities.TodoList) error {
// 	return s.repository.UpdateList(list)
// }

// func (s *TodoService) DeleteList(listId string) error {
// 	return s.repository.DeleteList(listId)
// }
