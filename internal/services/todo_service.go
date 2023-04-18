package services

import (
	"github.com/yotzapon/todo-service/internal/database"
	"github.com/yotzapon/todo-service/internal/entity"
)

type TodoServiceInterface interface {
	CreateTodo(input entity.TodoEntity) (*entity.TodoEntity, error)
	GetTodo(filter map[string]interface{}, order map[string]interface{}, limit int) ([]entity.TodoEntity, error)
	UpdateTodo(input entity.TodoEntity) (*entity.TodoEntity, error)
	MarkCompleteTodo(input entity.TodoEntity) (*entity.TodoEntity, error)
	DeleteTodo(input entity.TodoEntity) (*entity.TodoEntity, error)
}

type todoService struct {
	todoDB database.DB
}

func NewTodoService(db database.DB) TodoServiceInterface {
	return &todoService{todoDB: db}
}

func (t *todoService) CreateTodo(input entity.TodoEntity) (*entity.TodoEntity, error) {
	return t.todoDB.Todo().Create(input)
}

func (t *todoService) GetTodo(filter map[string]interface{}, order map[string]interface{}, limit int) ([]entity.TodoEntity, error) {
	return t.todoDB.Todo().Find(filter, order, limit)
}

func (t *todoService) UpdateTodo(input entity.TodoEntity) (*entity.TodoEntity, error) {
	return t.todoDB.Todo().Update(input)
}

func (t *todoService) MarkCompleteTodo(input entity.TodoEntity) (*entity.TodoEntity, error) {
	return t.todoDB.Todo().MarkComplete(input)
}

func (t *todoService) DeleteTodo(input entity.TodoEntity) (*entity.TodoEntity, error) {
	return t.todoDB.Todo().Delete(input)
}
