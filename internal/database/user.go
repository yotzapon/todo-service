package database

import (
	"gorm.io/gorm"

	"github.com/yotzapon/todo-service/internal/entity"
)

type userRepo struct {
	gorm *gorm.DB
}

type UserRepositoryInterface interface {
	Find(input entity.AuthEntity) (*entity.AuthEntity, error)
}

func (d db) User() UserRepositoryInterface {
	return userRepo(d)
}

func (t userRepo) Find(input entity.AuthEntity) (*entity.AuthEntity, error) {
	userDBModel := userDBModel{}
	err := t.gorm.Where("username = ? AND password = ?", input.Username, input.Password).First(&userDBModel).Error

	return userDBModel.toEntity(), err
}
