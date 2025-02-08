package service

import (
	"github.com/senyabanana/todo-app/internal/entity"
	"github.com/senyabanana/todo-app/internal/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type TodoLists interface {
	Create(userId int, list entity.TodoList) (int, error)
	GetAll(userId int) ([]entity.TodoList, error)
	GetById(userId, listId int) (entity.TodoList, error)
	Update(userId, listId int, input entity.UpdateListInput) error
	Delete(userId, listId int) error
}

type TodoItems interface {
	Create(userId, listId int, input entity.TodoItem) (int, error)
	GetAll(userId, listId int) ([]entity.TodoItem, error)
	GetById(userId, itemId int) (entity.TodoItem, error)
	Update(userId, itemId int, input entity.UpdateItemInput) error
	Delete(userId, itemId int) error
}

type Service struct {
	Authorization
	TodoLists
	TodoItems
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoLists:     NewTodoListsService(repos.TodoLists),
		TodoItems:     NewTodoItemsService(repos.TodoItems, repos.TodoLists),
	}
}
