package config

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/jinzhu/configor"
)

type Config struct {
	DB         DatabaseConfig       `json:"db"`
	AppConfig  AppConfig            `json:"appConfig"`
	AuthConfig AuthenticationConfig `json:"authConfig"`
}

type AppConfig struct {
	Port    int    `json:"port" env:"CONFIG__APP_CONFIG__PORT" default:"8080"`
	Version string `json:"version" env:"APP__VERSION" default:"local"`
}

type DatabaseConfig struct {
	Host     string `json:"host" env:"CONFIG__DB__HOST" required:"true"`
	Port     int    `json:"port" env:"CONFIG__DB__PORT" default:"3306"`
	User     string `json:"user" env:"CONFIG__DB__USER" required:"true"`
	Password string `json:"password" env:"CONFIG__DB__PASSWORD" required:"true"`
	Name     string `json:"name" env:"CONFIG__DB__NAME" required:"true"`
}

type AuthenticationConfig struct {
	SecretKey      string `json:"secretKey" env:"CONFIG__AUTH__SECRET_KEY" default:"my_secret_key"`
	AccessTokenTTL int    `json:"AccessTokenTTL" env:"CONFIG__AUTH__SECRET_KEY" default:"3600"`
}

func LoadConfig() (*Config, error) {
	var config Config
	err := configor.
		New(&configor.Config{AutoReload: false}).
		Load(&config, fmt.Sprintf("%s/config.%s.json", getConfigLocation(), getEnv()))

	if err != nil {
		return nil, err
	}

	return &config, nil
}

func getConfigLocation() string {
	_, filename, _, _ := runtime.Caller(0)

	return path.Join(path.Dir(filename), "../config")
}

func getEnv() string {
	val := os.Getenv("APP_ENV")
	// todo: check our stage names and align with them
	switch strings.ToLower(val) {
	case "prod":
		return "prod"
	case "staging":
		return "staging"
	case "test":
		return "test"
	case "qa":
		return "qa"
	default:
		return "dev"
	}
}
