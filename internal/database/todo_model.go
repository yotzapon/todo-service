package database

import (
	"github.com/jinzhu/copier"

	"github.com/yotzapon/todo-service/internal/constants"
	"github.com/yotzapon/todo-service/internal/entity"
)

type todoDBModel struct {
	Model
	Title       string
	Description string
	IsCompleted bool `gorm:"column:completed"`
	UserID      int
}

func (t *todoDBModel) TableName() string {
	return constants.TodoTable
}

func (t *todoDBModel) toDBModel(input *entity.TodoEntity) *todoDBModel {
	_ = copier.Copy(t, input)

	return t
}

func (t *todoDBModel) toEntity() *entity.TodoEntity {
	output := &entity.TodoEntity{}
	_ = copier.Copy(output, t)

	return output
}
