package main

import (
	"log"

	"github.com/yotzapon/todo-service/http"
)

// @title Todo APIs
// @version 1.0
// @description This is a Todo API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8082
// @BasePath /v1
func main() {
	log.Println(http.StartServer())
}
