//go:build integration_test
// +build integration_test

package database_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yotzapon/todo-service/internal/database"
	"github.com/yotzapon/todo-service/internal/entity"
)

func TestUserRepo_Find(t *testing.T) {
	t.Run("happy to find user", func(t *testing.T) {
		tx := MustGetRawDB(t).Begin()
		defer tx.Rollback()
		db := database.SetDB(tx)

		result, err := db.User().Find(entity.AuthEntity{
			Username: "tester01",
			Password: "1111",
		})
		assert.NoError(t, err)
		assert.Equal(t, "tester01", result.Username)
		assert.Equal(t, "1111", result.Password)
	})
}
