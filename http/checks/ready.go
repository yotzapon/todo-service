package checks

import (
	"net/http"

	"github.com/yotzapon/todo-service/config"
	"github.com/yotzapon/todo-service/internal/database"
)

// Readiness    	Readiness check
// @Summary      	This API is used to check the readiness of an application or service.
// @Description  	This API is used to check the readiness of an application or service.
// @Tags         	HealthCheck
// @Accept       	json
// @Produce      	json
// @Success      	200
// @Router       	/readyz [get]
func Readiness(db database.DB, config config.Config) http.HandlerFunc {
	check := newCheck(db, config)
	check.mustRegister(check.db())

	return check.handler()
}
