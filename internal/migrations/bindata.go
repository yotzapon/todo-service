//go:build bindata
// +build bindata

package migrations

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"gorm.io/gorm"

	data "github.com/yotzapon/todo-service/internal/bindata/migrations"
)

func New(db *gorm.DB, noLock bool) (*migrate.Migrate, error) {
	assets := bindata.Resource(data.AssetNames(), func(name string) ([]byte, error) {
		fmt.Println(name)
		return data.Asset(name)
	})
	driver, err := bindata.WithInstance(assets)
	if err != nil {
		return nil, err
	}
	migrationDriver, err := getDbDriver(db, noLock)
	if err != nil {
		return nil, err
	}
	return migrate.NewWithInstance("go-bindata", driver, "mysql", migrationDriver)
}
