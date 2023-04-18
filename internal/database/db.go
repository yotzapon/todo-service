package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/yotzapon/todo-service/config"
	"github.com/yotzapon/todo-service/internal/ulid"
)

type TransactionFunction func(DB) error

type db struct {
	gorm *gorm.DB
}

type DB interface {
	RawDB() *gorm.DB
	Todo() TodoRepositoryInterface
	User() UserRepositoryInterface
	Ping() error
}

func (d db) RawDB() *gorm.DB {
	return d.gorm
}

type Model struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (d *Model) BeforeUpdate(g *gorm.DB) error {
	d.UpdatedAt = time.Now()
	g.Statement.Omit("created_at")

	return nil
}

func (d *Model) BeforeCreate(g *gorm.DB) error {
	name := g.Statement.Table
	if d.ID != "" {
		return nil
	}
	d.ID = ulid.New(name)

	return nil
}

func SetDB(d *gorm.DB) DB {
	return db{gorm: d}
}

func New(dbConfig config.DatabaseConfig) (DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	return db{
		gorm: gormDb,
	}, nil
}
