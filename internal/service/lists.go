package service

import (
	"github.com/senyabanana/todo-app/internal/entity"
	"github.com/senyabanana/todo-app/internal/repository"
)

type TodoListsService struct {
	repo repository.TodoLists
}

func NewTodoListsService(repo repository.TodoLists) *TodoListsService {
	return &TodoListsService{repo: repo}
}

func (s *TodoListsService) Create(userId int, list entity.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListsService) GetAll(userId int) ([]entity.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListsService) GetById(userId, listId int) (entity.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListsService) Update(userId, listId int, input entity.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, listId, input)
}

func (s *TodoListsService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}
