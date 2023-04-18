//go:build integration_test
// +build integration_test

package database_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/yotzapon/todo-service/config"
	"github.com/yotzapon/todo-service/internal/database"
)

func MustGetRawDB(t *testing.T) *gorm.DB {
	cfg, _ := config.LoadConfig()
	db, err := database.New(cfg.DB)
	if err != nil {
		t.Failed()
	}

	return db.RawDB()
}

func TestDb_Ping(t *testing.T) {
	tx := MustGetRawDB(t)
	db := database.SetDB(tx)
	err := db.Ping()
	assert.NoError(t, err)
}
