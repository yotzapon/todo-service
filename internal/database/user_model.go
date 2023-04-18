package database

import (
	"github.com/jinzhu/copier"

	"github.com/yotzapon/todo-service/internal/constants"
	"github.com/yotzapon/todo-service/internal/entity"
)

type userDBModel struct {
	ID       int
	Username string
	Password string
}

func (u *userDBModel) TableName() string {
	return constants.UserTable
}

func (u *userDBModel) toEntity() *entity.AuthEntity {
	output := &entity.AuthEntity{}
	_ = copier.Copy(output, u)
	output.UserID = u.ID

	return output
}
