package auth

import "github.com/yotzapon/todo-service/internal/entity"

type RequestAuthModel struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *RequestAuthModel) toAuthEntity() entity.AuthEntity {
	return entity.AuthEntity{
		Username: r.Username,
		Password: r.Password,
	}
}

type ResponseAuthModel struct {
	AccessToken string `json:"access_token"`
}
