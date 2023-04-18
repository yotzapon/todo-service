//go:build integration_test
// +build integration_test

package database_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/yotzapon/todo-service/internal/database"
	"github.com/yotzapon/todo-service/internal/entity"
)

var titleTooLong = "test_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todotest_todo"

func TestTodoRepo_Create(t *testing.T) {
	t.Run("happy create todo", func(t *testing.T) {
		tx := MustGetRawDB(t).Begin()
		defer tx.Rollback()
		db := database.SetDB(tx)

		entity, err := db.Todo().Create(entity.TodoEntity{
			Title:       "test_todo",
			Description: "desc_todo",
			IsCompleted: false,
			UserID:      1,
		})

		assert.NoError(t, err)
		assert.Equal(t, 1, entity.UserID)
		assert.Equal(t, "desc_todo", entity.Description)
		assert.Equal(t, "test_todo", entity.Title)
	})

	t.Run("failed create todo: title too long", func(t *testing.T) {
		tx := MustGetRawDB(t).Begin()
		defer tx.Rollback()
		db := database.SetDB(tx)

		_, err := db.Todo().Create(entity.TodoEntity{
			Title:       titleTooLong,
			Description: "desc_todo",
			IsCompleted: false,
			UserID:      1,
		})

		assert.Error(t, err)
	})
}

func TestTodoRepo_Find(t *testing.T) {
	t.Run("happy to find todo", func(t *testing.T) {
		tx := MustGetRawDB(t).Begin()
		defer tx.Rollback()
		db := database.SetDB(tx)

		_, _ = db.Todo().Create(entity.TodoEntity{
			Title:       "test_todo",
			Description: "desc_todo",
			IsCompleted: false,
			UserID:      1,
		})

		filter := make(map[string]interface{})
		filter["user_id"] = 1

		result, err := db.Todo().Find(filter, nil, 0)
		assert.NoError(t, err)
		assert.Equal(t, 1, result[0].UserID)
		assert.Equal(t, "desc_todo", result[0].Description)
		assert.Equal(t, "test_todo", result[0].Title)
	})

	t.Run("happy to find todo with order and limit", func(t *testing.T) {
		tx := MustGetRawDB(t).Begin()
		defer tx.Rollback()
		db := database.SetDB(tx)

		_, _ = db.Todo().Create(entity.TodoEntity{
			Title:       "test_todo",
			Description: "desc_todo",
			IsCompleted: false,
			UserID:      1,
		})

		time.Sleep(time.Second)
		_, _ = db.Todo().Create(entity.TodoEntity{
			Title:       "test_todo_1",
			Description: "desc_todo_1",
			IsCompleted: false,
			UserID:      1,
		})

		filter := make(map[string]interface{})
		filter["user_id"] = 1

		order := make(map[string]interface{})
		order["created_at"] = "DESC"

		result, err := db.Todo().Find(filter, order, 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(result))
		assert.Equal(t, 1, result[0].UserID)
		assert.Equal(t, "desc_todo_1", result[0].Description)
		assert.Equal(t, "test_todo_1", result[0].Title)
	})

	t.Run("failed to find todo", func(t *testing.T) {
		tx := MustGetRawDB(t).Begin()
		defer tx.Rollback()
		db := database.SetDB(tx)

		_, _ = db.Todo().Create(entity.TodoEntity{
			Title:       "test_todo",
			Description: "desc_todo",
			IsCompleted: false,
			UserID:      1,
		})

		filter := make(map[string]interface{})
		filter["user_id"] = 2

		result, err := db.Todo().Find(filter, nil, 0)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(result))
	})
}

func TestTodoRepo_Update(t *testing.T) {
	t.Run("happy update todo", func(t *testing.T) {
		tx := MustGetRawDB(t).Begin()
		defer tx.Rollback()
		db := database.SetDB(tx)

		result, _ := db.Todo().Create(entity.TodoEntity{
			Title:       "test_todo",
			Description: "desc_todo",
			IsCompleted: false,
			UserID:      1,
		})

		rsUpdate, err := db.Todo().Update(entity.TodoEntity{
			ID:          result.ID,
			Title:       "test_todo_update",
			Description: "desc_todo_update",
			UserID:      1,
		})

		assert.NoError(t, err)
		assert.Equal(t, 1, rsUpdate.UserID)
		assert.Equal(t, "desc_todo_update", rsUpdate.Description)
		assert.Equal(t, "test_todo_update", rsUpdate.Title)
	})

	t.Run("failed to update: not found todo ID", func(t *testing.T) {
		tx := MustGetRawDB(t).Begin()
		defer tx.Rollback()
		db := database.SetDB(tx)

		_, err := db.Todo().Update(entity.TodoEntity{
			ID:          "1",
			Title:       "test_todo_update",
			Description: "desc_todo_update",
			UserID:      1,
		})

		assert.Error(t, err)
	})

	t.Run("failed to update: title too long", func(t *testing.T) {
		tx := MustGetRawDB(t).Begin()
		defer tx.Rollback()
		db := database.SetDB(tx)

		result, _ := db.Todo().Create(entity.TodoEntity{
			Title:       "test_todo",
			Description: "desc_todo",
			IsCompleted: false,
			UserID:      1,
		})

		_, err := db.Todo().Update(entity.TodoEntity{
			ID:          result.ID,
			Title:       titleTooLong,
			Description: "desc_todo_update",
			UserID:      1,
		})

		assert.Error(t, err)
	})
}

func TestTodoRepo_MarkComplete(t *testing.T) {
	t.Run("happy mark todo complete", func(t *testing.T) {
		tx := MustGetRawDB(t).Begin()
		defer tx.Rollback()
		db := database.SetDB(tx)

		result, _ := db.Todo().Create(entity.TodoEntity{
			Title:       "test_todo",
			Description: "desc_todo",
			IsCompleted: false,
			UserID:      1,
		})

		rsUpdate, err := db.Todo().MarkComplete(entity.TodoEntity{
			ID:          result.ID,
			IsCompleted: true,
			UserID:      1,
		})

		assert.NoError(t, err)
		assert.Equal(t, 1, rsUpdate.UserID)
		assert.Equal(t, "desc_todo", rsUpdate.Description)
		assert.Equal(t, "test_todo", rsUpdate.Title)
		assert.True(t, rsUpdate.IsCompleted)
	})

	t.Run("failed to mark todo complete", func(t *testing.T) {
		tx := MustGetRawDB(t).Begin()
		defer tx.Rollback()
		db := database.SetDB(tx)

		_, err := db.Todo().MarkComplete(entity.TodoEntity{
			ID:          "2",
			IsCompleted: true,
			UserID:      2,
		})

		assert.Error(t, err)
	})
}

func TestTodoRepo_Delete(t *testing.T) {
	t.Run("happy delete todo", func(t *testing.T) {
		tx := MustGetRawDB(t).Begin()
		defer tx.Rollback()
		db := database.SetDB(tx)

		result, _ := db.Todo().Create(entity.TodoEntity{
			Title:       "test_todo",
			Description: "desc_todo",
			IsCompleted: false,
			UserID:      1,
		})

		_, err := db.Todo().Delete(entity.TodoEntity{
			ID:     result.ID,
			UserID: 1,
		})
		assert.NoError(t, err)

		filter := make(map[string]interface{})
		filter["user_id"] = 1

		findResult, findErr := db.Todo().Find(filter, nil, 0)
		assert.NoError(t, findErr)
		assert.Equal(t, 0, len(findResult))
	})
}
