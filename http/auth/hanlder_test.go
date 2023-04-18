package auth_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"

	"github.com/yotzapon/todo-service/config"
	localHttp "github.com/yotzapon/todo-service/http"
	"github.com/yotzapon/todo-service/internal/entity"
	"github.com/yotzapon/todo-service/internal/services"
	"github.com/yotzapon/todo-service/mocks"
)

func TestHandler_Login(t *testing.T) {
	executeWithRequest := func(authServ services.AuthServiceInterface,
		todoServ services.TodoServiceInterface,
		cfg *config.Config, body string) *httptest.ResponseRecorder {
		response := httptest.NewRecorder()
		_, ginEngine := gin.CreateTestContext(response)

		requestURL := "/login"
		httpRequest, _ := http.NewRequest(http.MethodPost, requestURL, strings.NewReader(body))

		localHttp.NewRoutes(ginEngine, authServ, todoServ, cfg)
		ginEngine.ServeHTTP(response, httpRequest)

		return response
	}

	t.Run("happy login", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg, _ := config.LoadConfig()
		servAuth := mocks.NewMockAuthServiceInterface(ctrl)
		servTodo := mocks.NewMockTodoServiceInterface(ctrl)
		servAuth.EXPECT().Login(gomock.Any()).Return(mockAuthEntity(), nil)
		servAuth.EXPECT().GenerateJWT(gomock.Any()).Return("token", nil)
		resp := executeWithRequest(servAuth, servTodo, cfg, bodyRequest())
		b, _ := io.ReadAll(resp.Body)

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, "token", gjson.GetBytes(b, "access_token").String())
	})

	t.Run("failed invalid body request", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg, _ := config.LoadConfig()
		servAuth := mocks.NewMockAuthServiceInterface(ctrl)
		servTodo := mocks.NewMockTodoServiceInterface(ctrl)
		resp := executeWithRequest(servAuth, servTodo, cfg, invalidBodyRequest())

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("failed user not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg, _ := config.LoadConfig()
		servAuth := mocks.NewMockAuthServiceInterface(ctrl)
		servTodo := mocks.NewMockTodoServiceInterface(ctrl)
		servAuth.EXPECT().Login(gomock.Any()).Return(mockAuthEntity(), errors.New("should error"))
		resp := executeWithRequest(servAuth, servTodo, cfg, bodyRequest())

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusUnauthorized, resp.Code)
	})

	t.Run("failed generate token", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg, _ := config.LoadConfig()
		servAuth := mocks.NewMockAuthServiceInterface(ctrl)
		servTodo := mocks.NewMockTodoServiceInterface(ctrl)
		servAuth.EXPECT().Login(gomock.Any()).Return(mockAuthEntity(), nil)
		servAuth.EXPECT().GenerateJWT(gomock.Any()).Return("token", errors.New("should error"))
		resp := executeWithRequest(servAuth, servTodo, cfg, bodyRequest())

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusUnauthorized, resp.Code)
	})
}

func mockAuthEntity() *entity.AuthEntity {
	return &entity.AuthEntity{
		UserID:   1,
		Username: "tester01",
		Password: "1111",
	}
}

func bodyRequest() string {
	return `{
    "username":"tester01",
    "password":"1111"
}`
}

func invalidBodyRequest() string {
	return `{
    username:"tester01",
    "password":"1111"
}`
}
