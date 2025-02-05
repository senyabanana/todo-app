package service

import "github.com/senyabanana/todo-app/internal/repository"

type Authorization interface {
}

type TodoLists interface {
}

type TodoItems interface {
}

type Service struct {
	Authorization
	TodoLists
	TodoItems
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
