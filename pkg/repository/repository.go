package repository

import (
	"github.com/Yujiman/GoTodo"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoItem interface {
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {

	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoLstPostgre(db),
	}
}
