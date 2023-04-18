package todo

import (
	"time"

	"github.com/jinzhu/copier"

	"github.com/yotzapon/todo-service/internal/entity"
)

type RequestTodoModel struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type ResponseTodoModel struct {
	ID          string    `json:"id" example:"todos_01GY7687DW357TEAQR5TA23VEA"`
	Title       string    `json:"title" example:"home"`
	Description string    `json:"description" example:"buy the bag"`
	IsCompleted bool      `json:"isCompleted" example:"false"`
	CreatedAt   time.Time `json:"created" example:"2023-04-17T15:45:39+07:00"`
	UpdatedAt   time.Time `json:"updated" example:"2023-04-17T15:45:39+07:00"`
}

func toTodoResponse(input *entity.TodoEntity) ResponseTodoModel {
	resp := &ResponseTodoModel{}
	_ = copier.Copy(resp, input)

	return *resp
}

func toTodoResponses(inputs []entity.TodoEntity) []ResponseTodoModel {
	resp := []ResponseTodoModel{}

	item := ResponseTodoModel{}
	for _, v := range inputs {
		_ = copier.Copy(&item, &v)

		resp = append(resp, item)
	}

	return resp
}

func (r *RequestTodoModel) toTodoEntity() entity.TodoEntity {
	return entity.TodoEntity{
		Title:       r.Title,
		Description: r.Description,
	}
}
