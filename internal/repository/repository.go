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
}

type TodoItems interface {
}

type Repository struct {
	Authorization
	TodoLists
	TodoItems
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
