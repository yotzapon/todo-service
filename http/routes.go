package http

import (
	"github.com/gin-gonic/gin"

	"github.com/yotzapon/todo-service/config"
	"github.com/yotzapon/todo-service/http/auth"
	"github.com/yotzapon/todo-service/http/middlewares"
	"github.com/yotzapon/todo-service/http/todo"
	"github.com/yotzapon/todo-service/internal/services"
)

func NewRoutes(ginEngine *gin.Engine,
	servAuth services.AuthServiceInterface,
	servTodo services.TodoServiceInterface,
	cfg *config.Config) {
	ginEngine.POST("/login", auth.NewAuthHandler(servAuth).Login)
	authorized := ginEngine.Group("v1")
	authorized.Use(middlewares.AuthMiddleware(cfg))
	{
		authorized.POST("/todos", todo.NewTodoHandler(servTodo, cfg).CreateTodo)
		authorized.PUT("/todos", todo.NewTodoHandler(servTodo, cfg).UpdateTodo)
		authorized.PATCH("/todos", todo.NewTodoHandler(servTodo, cfg).PatchTodo)
		authorized.DELETE("/todos", todo.NewTodoHandler(servTodo, cfg).DeleteTodo)
		authorized.GET("/todos", todo.NewTodoHandler(servTodo, cfg).GetTodo)
	}
}
