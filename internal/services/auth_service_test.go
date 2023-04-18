package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/yotzapon/todo-service/config"
	"github.com/yotzapon/todo-service/internal/entity"
	"github.com/yotzapon/todo-service/internal/services"
	"github.com/yotzapon/todo-service/mocks"
)

func TestAuthService(t *testing.T) {
	t.Run("happy login", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		db := mocks.NewMockDB(ctrl)
		userRepo := mocks.NewMockUserRepositoryInterface(ctrl)
		db.EXPECT().User().Return(userRepo)
		userRepo.EXPECT().Find(gomock.Any()).Return(&entity.AuthEntity{Username: "tester01"}, nil)
		cfg, _ := config.LoadConfig()

		serv := services.NewAuthService(db, cfg)
		auth, err := serv.Login(entity.AuthEntity{})

		assert.NoError(t, err)
		assert.Equal(t, "tester01", auth.Username)
	})

	t.Run("happy generate JWT", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		db := mocks.NewMockDB(ctrl)
		cfg, _ := config.LoadConfig()

		serv := services.NewAuthService(db, cfg)
		token, err := serv.GenerateJWT(entity.AuthEntity{Username: "tester01", UserID: 1})

		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})
}
