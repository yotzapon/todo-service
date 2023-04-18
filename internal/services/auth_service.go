package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/yotzapon/todo-service/config"
	"github.com/yotzapon/todo-service/internal/database"
	"github.com/yotzapon/todo-service/internal/entity"
)

type AuthServiceInterface interface {
	Login(input entity.AuthEntity) (*entity.AuthEntity, error)
	GenerateJWT(input entity.AuthEntity) (string, error)
}

type authService struct {
	authDB database.DB
	config *config.Config
}

type Claims struct {
	Username string `json:"username"`
	Id       int    `json:"id"`
	jwt.StandardClaims
}

func NewAuthService(db database.DB, cfg *config.Config) AuthServiceInterface {
	return &authService{authDB: db, config: cfg}
}

func (a *authService) Login(input entity.AuthEntity) (*entity.AuthEntity, error) {
	return a.authDB.User().Find(input)
}

func (a *authService) GenerateJWT(input entity.AuthEntity) (string, error) {
	// Create a JWT token with a 24-hour expiration time
	claims := &Claims{
		Username: input.Username,
		Id:       input.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(a.config.AuthConfig.AccessTokenTTL)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(a.config.AuthConfig.SecretKey))
}
