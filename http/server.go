package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/yotzapon/todo-service/config"
	_ "github.com/yotzapon/todo-service/docs"
	"github.com/yotzapon/todo-service/http/checks"
	"github.com/yotzapon/todo-service/internal/database"
	"github.com/yotzapon/todo-service/internal/services"
	"github.com/yotzapon/todo-service/internal/xlogger"
)

var logger = xlogger.Get()

func StartServer() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ginEngine := gin.New()
	ginEngine.Use(gin.Logger())
	ginEngine.Use(gin.Recovery())
	ginEngine.Use(
		requestid.New(
			requestid.WithGenerator(func() string {
				return uuid.NewString()
			}),
			requestid.WithCustomHeaderStrKey("x-request-id"),
		),
	)

	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8082/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1)))

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := database.New(cfg.DB)
	if err != nil {
		panic(err)
	}

	ginEngine.GET("/livez", gin.WrapF(checks.Liveness(db, *cfg)))
	ginEngine.GET("/readyz", gin.WrapF(checks.Readiness(db, *cfg)))

	// init services
	servTodo := services.NewTodoService(db)
	servAuth := services.NewAuthService(db, cfg)

	NewRoutes(ginEngine, servAuth, servTodo, cfg)

	port := fmt.Sprintf(":%v", cfg.AppConfig.Port)
	server := &http.Server{
		Addr:    port,
		Handler: ginEngine,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}
	}()

	logger.Infof("server ready to serve port: %v\n", port)
	// Wait for interrupt signal to gracefully shutdown the ginEngine with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down ginEngine...")

	// The context is used to inform the ginEngine it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown: ", err)
	}

	logger.Println("Server exiting")

	return nil
}
