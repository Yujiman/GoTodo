package service

import (
	todo "github.com/Yujiman/GoTodo"
	"github.com/Yujiman/GoTodo/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string)(int, error)
}

type TodoItem interface {
}

type TodoList interface {
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
