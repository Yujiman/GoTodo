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

func (t *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return t.repo.GetAll(userId)
}

func (t *TodoListService) GetById(userId, listId int) (todo.TodoList, error) {
	return t.repo.GetById(userId, listId)
}

func (t *TodoListService) Delete(userId, listId int) error {
	return t.repo.Delete(userId, listId)

}

func (t *TodoListService) Update(userId, listId int, input todo.UpdatedListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return t.repo.Update(userId, listId, input)
}
