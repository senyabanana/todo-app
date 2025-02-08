package service

import (
	"github.com/senyabanana/todo-app/internal/entity"
	"github.com/senyabanana/todo-app/internal/repository"
)

type TodoItemsService struct {
	repo     repository.TodoItems
	listRepo repository.TodoLists
}

func NewTodoItemsService(repo repository.TodoItems, listRepo repository.TodoLists) *TodoItemsService {
	return &TodoItemsService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemsService) Create(userId, listId int, input entity.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId, input)
}

func (s *TodoItemsService) GetAll(userId, listId int) ([]entity.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemsService) GetById(userId, itemId int) (entity.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemsService) Update(userId, itemId int, input entity.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, itemId, input)
}

func (s *TodoItemsService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}
