package checks

import (
	"net/http"

	"github.com/yotzapon/todo-service/config"
	"github.com/yotzapon/todo-service/internal/database"
)

// Liveness    		Liveness check
// @Summary      	This API is used to check the liveness of an application or service.
// @Description  	This API is used to check the liveness of an application or service.
// @Tags         	HealthCheck
// @Accept       	json
// @Produce      	json
// @Success      	200
// @Router       	/livez [get]
func Liveness(db database.DB, config config.Config) http.HandlerFunc {
	check := newCheck(db, config)
	check.mustRegister(check.db())

	return check.handler()
}
