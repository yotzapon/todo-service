//go:build !bindata
// +build !bindata

package migrations

import (
	"path"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

func New(db *gorm.DB, noLock bool) (*migrate.Migrate, error) {
	dbDriver, err := getDbDriver(db, noLock)
	if err != nil {
		return nil, err
	}
	_, filename, _, _ := runtime.Caller(0)
	migrationsPath := path.Join(path.Dir(filename), "../../migrations")

	return migrate.NewWithDatabaseInstance("file:///"+migrationsPath, "postgres", dbDriver)
}
