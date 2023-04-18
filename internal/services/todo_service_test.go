package services_test

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/yotzapon/todo-service/internal/entity"
	"github.com/yotzapon/todo-service/internal/services"
	"github.com/yotzapon/todo-service/mocks"
)

func TestTodoService(t *testing.T) {
	t.Run("happy create", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		db := mocks.NewMockDB(ctrl)
		todoRepo := mocks.NewMockTodoRepositoryInterface(ctrl)
		db.EXPECT().Todo().Return(todoRepo)

		todoRepo.EXPECT().Create(gomock.Any()).Return(mockTodoEntity(), nil)

		serv := services.NewTodoService(db)
		todo, err := serv.CreateTodo(entity.TodoEntity{})

		assert.NoError(t, err)
		assert.Equal(t, "test_todo", todo.Title)
	})

	t.Run("happy get todo", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		db := mocks.NewMockDB(ctrl)
		todoRepo := mocks.NewMockTodoRepositoryInterface(ctrl)
		db.EXPECT().Todo().Return(todoRepo)

		todoRepo.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockTodosEntity(), nil)

		serv := services.NewTodoService(db)
		todo, err := serv.GetTodo(nil, nil, 0)

		assert.NoError(t, err)
		assert.Equal(t, "test_todo", todo[0].Title)
	})

	t.Run("happy update todo", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		db := mocks.NewMockDB(ctrl)
		todoRepo := mocks.NewMockTodoRepositoryInterface(ctrl)
		db.EXPECT().Todo().Return(todoRepo)

		todoRepo.EXPECT().Update(gomock.Any()).Return(mockTodoEntity(), nil)

		serv := services.NewTodoService(db)
		todo, err := serv.UpdateTodo(entity.TodoEntity{})

		assert.NoError(t, err)
		assert.Equal(t, "test_todo", todo.Title)
	})

	t.Run("happy mark todo complete", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		db := mocks.NewMockDB(ctrl)
		todoRepo := mocks.NewMockTodoRepositoryInterface(ctrl)
		db.EXPECT().Todo().Return(todoRepo)

		todoRepo.EXPECT().MarkComplete(gomock.Any()).Return(mockTodoEntity(), nil)

		serv := services.NewTodoService(db)
		todo, err := serv.MarkCompleteTodo(entity.TodoEntity{})

		assert.NoError(t, err)
		assert.Equal(t, "test_todo", todo.Title)
	})

	t.Run("happy delete todo", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		db := mocks.NewMockDB(ctrl)
		todoRepo := mocks.NewMockTodoRepositoryInterface(ctrl)
		db.EXPECT().Todo().Return(todoRepo)

		todoRepo.EXPECT().Delete(gomock.Any()).Return(mockTodoEntity(), nil)

		serv := services.NewTodoService(db)
		todo, err := serv.DeleteTodo(entity.TodoEntity{})

		assert.NoError(t, err)
		assert.Equal(t, "test_todo", todo.Title)
	})
}

func mockTodoEntity() *entity.TodoEntity {
	return &entity.TodoEntity{
		ID:          "0001",
		Title:       "test_todo",
		Description: "desc_todo",
		IsCompleted: false,
		UserID:      1,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}

func mockTodosEntity() []entity.TodoEntity {
	return []entity.TodoEntity{*mockTodoEntity()}
}
