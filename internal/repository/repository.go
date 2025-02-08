package repository

import (
	"github.com/senyabanana/todo-app/internal/entity"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, password string) (entity.User, error)
}

type TodoLists interface {
	Create(userId int, list entity.TodoList) (int, error)
	GetAll(userId int) ([]entity.TodoList, error)
	GetById(userId, listId int) (entity.TodoList, error)
	Update(userId, listId int, input entity.UpdateListInput) error
	Delete(userId, listId int) error
}

type TodoItems interface {
	Create(listId int, input entity.TodoItem) (int, error)
	GetAll(userId, listId int) ([]entity.TodoItem, error)
	GetById(userId, itemId int) (entity.TodoItem, error)
	Update(userId, itemId int, input entity.UpdateItemInput) error
	Delete(userId, itemId int) error
}

type Repository struct {
	Authorization
	TodoLists
	TodoItems
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoLists:     NewTodoListsPostgres(db),
		TodoItems:     NewTodoItemsRepository(db),
	}
}
