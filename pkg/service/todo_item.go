package service

import (
	"github.com/zhashkevych/todo-app"
	"github.com/zhashkevych/todo-app/pkg/repository"
)

type todoItemService struct {
	repo     repository.TodoItem
	listrepo repository.TodoList
}

func newTodoItemSerivce(repo repository.TodoItem, listrepo repository.TodoList) *todoItemService {
	return &todoItemService{repo: repo, listrepo: listrepo}
}

func (s *todoItemService) Create(userId int, listId int, item todo.TodoItem) (int, error) {
	_, err := s.listrepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *todoItemService) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *todoItemService) GetById(userId, itemId int) (todo.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *todoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *todoItemService) Update(userId, itemId int, input todo.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}
