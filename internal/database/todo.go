package database

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/yotzapon/todo-service/internal/entity"
	"github.com/yotzapon/todo-service/internal/xlogger"
)

var log = xlogger.Get()

type todoRepo struct {
	gorm *gorm.DB
}

type TodoRepositoryInterface interface {
	Create(input entity.TodoEntity) (*entity.TodoEntity, error)
	Find(filter map[string]interface{}, order map[string]interface{}, limit int) ([]entity.TodoEntity, error)
	Update(input entity.TodoEntity) (*entity.TodoEntity, error)
	MarkComplete(input entity.TodoEntity) (*entity.TodoEntity, error)
	Delete(input entity.TodoEntity) (*entity.TodoEntity, error)
}

func (d db) Todo() TodoRepositoryInterface {
	return todoRepo(d)
}

func (t todoRepo) Create(input entity.TodoEntity) (*entity.TodoEntity, error) {
	todoDBModel := todoDBModel{}
	model := todoDBModel.toDBModel(&input)
	err := t.gorm.Create(model).Error

	return model.toEntity(), err
}

func (t todoRepo) Find(filter map[string]interface{}, order map[string]interface{}, limit int) ([]entity.TodoEntity, error) {
	var todos []todoDBModel
	var result []entity.TodoEntity
	var err error

	where := t.gorm.Where(filter)

	for k, o := range order {
		log.Infof("key: %v, order: %v", k, o)
		where.Order(fmt.Sprintf("%v %v", k, o))
	}

	if limit > 0 {
		where.Limit(limit)
	}

	err = where.Find(&todos).Error
	if err != nil {
		return nil, err
	}

	for _, v := range todos {
		result = append(result, entity.TodoEntity{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			IsCompleted: v.IsCompleted,
			UserID:      v.UserID,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}

	return result, err
}

func (t todoRepo) Update(input entity.TodoEntity) (*entity.TodoEntity, error) {
	model := todoDBModel{}
	err := t.gorm.Where("id = ?", input.ID).First(&model).Error
	if err != nil {
		return nil, err
	}

	model.Title = input.Title
	model.Description = input.Description
	err = t.gorm.Where("user_id = ?", input.UserID).Updates(&model).Error

	return model.toEntity(), err
}

func (t todoRepo) MarkComplete(input entity.TodoEntity) (*entity.TodoEntity, error) {
	model := todoDBModel{}
	err := t.gorm.Where("id = ?", input.ID).First(&model).Error
	if err != nil {
		return nil, err
	}

	model.IsCompleted = input.IsCompleted
	err = t.gorm.Where("user_id = ?", input.UserID).Updates(&model).Error

	return model.toEntity(), err
}

func (t todoRepo) Delete(input entity.TodoEntity) (*entity.TodoEntity, error) {
	model := todoDBModel{}
	err := t.gorm.Where("id = ? AND user_id = ?", input.ID, input.UserID).Delete(&model).Error

	return model.toEntity(), err
}
