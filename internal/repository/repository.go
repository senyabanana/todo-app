package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/senyabanana/todo-app/internal/entity"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
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
