package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
	return &Repository{}
}
