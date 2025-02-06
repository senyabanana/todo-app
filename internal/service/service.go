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
}

type TodoItems interface {
}

type Service struct {
	Authorization
	TodoLists
	TodoItems
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
