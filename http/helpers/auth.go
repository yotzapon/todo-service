package helpers

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"

	"github.com/yotzapon/todo-service/config"
	"github.com/yotzapon/todo-service/internal/entity"
	"github.com/yotzapon/todo-service/internal/xlogger"
)

var logger = xlogger.Get()

func DecodeAccessToken(accessToken string, cfg *config.Config) *entity.AuthEntity {
	tokenString := accessToken[len("Bearer "):]

	// Parse the JWT token string into a token object
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method used
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key used to sign the token
		return []byte(cfg.AuthConfig.SecretKey), nil
	})
	if err != nil {
		logger.Error("failed to parse token:", err)

		return nil
	}

	// Check if the token is valid
	if !token.Valid {
		logger.Error("token is invalid")

		return nil
	}

	// Extract the claims from the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &entity.AuthEntity{
			UserID:   int(claims["id"].(float64)),
			Username: claims["username"].(string),
		}
	}

	return nil
}
