package todo_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"

	"github.com/yotzapon/todo-service/config"
	localHttp "github.com/yotzapon/todo-service/http"
	"github.com/yotzapon/todo-service/http/middlewares"
	"github.com/yotzapon/todo-service/internal/entity"
	"github.com/yotzapon/todo-service/internal/services"
	"github.com/yotzapon/todo-service/mocks"
)

var requestURL = "/v1/todos"

func TestHandler_Unauthorized(t *testing.T) {
	executeWithRequest := func(authServ services.AuthServiceInterface,
		todoServ services.TodoServiceInterface,
		cfg *config.Config, body string) *httptest.ResponseRecorder {
		response := httptest.NewRecorder()
		_, ginEngine := gin.CreateTestContext(response)

		ginEngine.Use(middlewares.AuthMiddleware(cfg))

		requestURL := "/v1/todos"
		httpRequest, _ := http.NewRequest(http.MethodPost, requestURL, strings.NewReader(body))

		localHttp.NewRoutes(ginEngine, authServ, todoServ, cfg)
		ginEngine.ServeHTTP(response, httpRequest)

		return response
	}

	t.Run("failed create todo unauthorized", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg, _ := config.LoadConfig()
		servAuth := mocks.NewMockAuthServiceInterface(ctrl)
		servTodo := mocks.NewMockTodoServiceInterface(ctrl)
		resp := executeWithRequest(servAuth, servTodo, cfg, bodyRequest())

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusUnauthorized, resp.Code)
	})
}

func TestHandler(t *testing.T) {
	executeWithRequest := func(authServ services.AuthServiceInterface,
		todoServ services.TodoServiceInterface,
		cfg *config.Config, method string, body string, requestURL string) *httptest.ResponseRecorder {
		response := httptest.NewRecorder()
		_, ginEngine := gin.CreateTestContext(response)

		ginEngine.Use(middlewares.AuthMiddleware(cfg))
		httpRequest, _ := http.NewRequest(method, requestURL, strings.NewReader(body))

		serv := services.NewAuthService(nil, cfg)
		token, _ := serv.GenerateJWT(entity.AuthEntity{UserID: 1, Username: "tester01"})
		httpRequest.Header.Add("Authorization", "Bearer "+token)
		localHttp.NewRoutes(ginEngine, authServ, todoServ, cfg)
		ginEngine.ServeHTTP(response, httpRequest)

		return response
	}

	t.Run("happy create todo", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg, _ := config.LoadConfig()
		servAuth := mocks.NewMockAuthServiceInterface(ctrl)
		servTodo := mocks.NewMockTodoServiceInterface(ctrl)
		servTodo.EXPECT().CreateTodo(gomock.Any()).Return(mockTodoEntity(), nil)
		resp := executeWithRequest(servAuth, servTodo, cfg, http.MethodPost, bodyRequest(), requestURL)
		b, _ := io.ReadAll(resp.Body)

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusCreated, resp.Code)
		assert.Equal(t, "home", gjson.GetBytes(b, "data.title").String())
	})

	t.Run("failed invalid body request", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg, _ := config.LoadConfig()
		servAuth := mocks.NewMockAuthServiceInterface(ctrl)
		servTodo := mocks.NewMockTodoServiceInterface(ctrl)
		resp := executeWithRequest(servAuth, servTodo, cfg, http.MethodPost, invalidBodyRequest(), requestURL)

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("failed create todo", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg, _ := config.LoadConfig()
		servAuth := mocks.NewMockAuthServiceInterface(ctrl)
		servTodo := mocks.NewMockTodoServiceInterface(ctrl)
		servTodo.EXPECT().CreateTodo(gomock.Any()).Return(mockTodoEntity(), errors.New("should error"))
		resp := executeWithRequest(servAuth, servTodo, cfg, http.MethodPost, bodyRequest(), requestURL)

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})

	t.Run("happy get todo", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg, _ := config.LoadConfig()
		servAuth := mocks.NewMockAuthServiceInterface(ctrl)
		servTodo := mocks.NewMockTodoServiceInterface(ctrl)
		servTodo.EXPECT().GetTodo(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockTodosEntity(), nil)
		resp := executeWithRequest(servAuth, servTodo, cfg, http.MethodGet, bodyRequest(), requestURL+"?ids=1,2&isComplete=false&orderCreated=desc&orderUpdated=asc&limit=2")
		b, _ := io.ReadAll(resp.Body)

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, "[\"home\"]", gjson.GetBytes(b, "data.#.title").String())
	})

	t.Run("failed get todo", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg, _ := config.LoadConfig()
		servAuth := mocks.NewMockAuthServiceInterface(ctrl)
		servTodo := mocks.NewMockTodoServiceInterface(ctrl)
		servTodo.EXPECT().GetTodo(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockTodosEntity(), errors.New("should error"))
		resp := executeWithRequest(servAuth, servTodo, cfg, http.MethodGet, bodyRequest(), requestURL+"?ids=1,2&isComplete=false&orderCreated=desc&limit=2")

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})

	t.Run("failed parse isComplete query param", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg, _ := config.LoadConfig()
		servAuth := mocks.NewMockAuthServiceInterface(ctrl)
		servTodo := mocks.NewMockTodoServiceInterface(ctrl)
		resp := executeWithRequest(servAuth, servTodo, cfg, http.MethodGet, bodyRequest(), requestURL+"?ids=1,2&isComplete=invalid&orderCreated=desc&limit=2")

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("failed parse limit query param", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg, _ := config.LoadConfig()
		servAuth := mocks.NewMockAuthServiceInterface(ctrl)
		servTodo := mocks.NewMockTodoServiceInterface(ctrl)
		resp := executeWithRequest(servAuth, servTodo, cfg, http.MethodGet, bodyRequest(), requestURL+"?ids=1,2&isComplete=false&orderCreated=desc&limit=a")

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("happy update todo", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg, _ := config.LoadConfig()
		servAuth := mocks.NewMockAuthServiceInterface(ctrl)
		servTodo := mocks.NewMockTodoServiceInterface(ctrl)
		servTodo.EXPECT().UpdateTodo(gomock.Any()).Return(mockTodoEntity(), nil)
		resp := executeWithRequest(servAuth, servTodo, cfg, http.MethodPut, bodyRequest(), requestURL)
		b, _ := io.ReadAll(resp.Body)

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, "home", gjson.GetBytes(b, "data.title").String())
	})

	t.Run("failed update todo", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg, _ := config.LoadConfig()
		servAuth := mocks.NewMockAuthServiceInterface(ctrl)
		servTodo := mocks.NewMockTodoServiceInterface(ctrl)
		servTodo.EXPECT().UpdateTodo(gomock.Any()).Return(mockTodoEntity(), errors.New("should error"))
		resp := executeWithRequest(servAuth, servTodo, cfg, http.MethodPut, bodyRequest(), requestURL)

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("failed update body request", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg, _ := config.LoadConfig()
		servAuth := mocks.NewMockAuthServiceInterface(ctrl)
		servTodo := mocks.NewMockTodoServiceInterface(ctrl)
		resp := executeWithRequest(servAuth, servTodo, cfg, http.MethodPut, invalidBodyRequest(), requestURL)

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("happy patch todo", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg, _ := config.LoadConfig()
		servAuth := mocks.NewMockAuthServiceInterface(ctrl)
		servTodo := mocks.NewMockTodoServiceInterface(ctrl)
		servTodo.EXPECT().MarkCompleteTodo(gomock.Any()).Return(mockTodoEntity(), nil)
		resp := executeWithRequest(servAuth, servTodo, cfg, http.MethodPatch, bodyRequest(), requestURL+"?id=1")
		b, _ := io.ReadAll(resp.Body)

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, "home", gjson.GetBytes(b, "data.title").String())
	})

	t.Run("failed patch todo", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg, _ := config.LoadConfig()
		servAuth := mocks.NewMockAuthServiceInterface(ctrl)
		servTodo := mocks.NewMockTodoServiceInterface(ctrl)
		servTodo.EXPECT().MarkCompleteTodo(gomock.Any()).Return(mockTodoEntity(), errors.New("should error"))
		resp := executeWithRequest(servAuth, servTodo, cfg, http.MethodPatch, bodyRequest(), requestURL+"?id=1")

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("failed patch todo missing query param", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg, _ := config.LoadConfig()
		servAuth := mocks.NewMockAuthServiceInterface(ctrl)
		servTodo := mocks.NewMockTodoServiceInterface(ctrl)
		resp := executeWithRequest(servAuth, servTodo, cfg, http.MethodPatch, bodyRequest(), requestURL)

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("happy delete todo", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg, _ := config.LoadConfig()
		servAuth := mocks.NewMockAuthServiceInterface(ctrl)
		servTodo := mocks.NewMockTodoServiceInterface(ctrl)
		servTodo.EXPECT().DeleteTodo(gomock.Any()).Return(mockTodoEntity(), nil)
		resp := executeWithRequest(servAuth, servTodo, cfg, http.MethodDelete, bodyRequest(), requestURL+"?id=1")

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusNoContent, resp.Code)
	})

	t.Run("failed delete todo", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg, _ := config.LoadConfig()
		servAuth := mocks.NewMockAuthServiceInterface(ctrl)
		servTodo := mocks.NewMockTodoServiceInterface(ctrl)
		servTodo.EXPECT().DeleteTodo(gomock.Any()).Return(mockTodoEntity(), errors.New("should error"))
		resp := executeWithRequest(servAuth, servTodo, cfg, http.MethodDelete, bodyRequest(), requestURL+"?id=1")

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})

	t.Run("failed delete todo missing required param", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg, _ := config.LoadConfig()
		servAuth := mocks.NewMockAuthServiceInterface(ctrl)
		servTodo := mocks.NewMockTodoServiceInterface(ctrl)
		resp := executeWithRequest(servAuth, servTodo, cfg, http.MethodDelete, bodyRequest(), requestURL)

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})
}

func bodyRequest() string {
	return `{
		"title":"home",
		"description":"buy the bag"
	}`
}

func invalidBodyRequest() string {
	return `{
		title:"home",
		"description":"buy the bag"
	}`
}

func mockTodoEntity() *entity.TodoEntity {
	return &entity.TodoEntity{
		ID:          "0001",
		Title:       "home",
		Description: "buy the bag",
		IsCompleted: false,
		UserID:      1,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}

func mockTodosEntity() []entity.TodoEntity {
	return []entity.TodoEntity{*mockTodoEntity()}
}
