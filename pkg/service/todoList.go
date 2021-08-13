package service

import (
	todo "github.com/Yujiman/GoTodo"
	"github.com/Yujiman/GoTodo/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodolistService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (t *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return t.repo.Create(userId, list)
}
