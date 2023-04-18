package checks

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hellofresh/health-go/v4"

	"github.com/yotzapon/todo-service/config"
	"github.com/yotzapon/todo-service/internal/database"
)

type check struct {
	_db     database.DB
	config  config.Config
	_health *health.Health
}

func newCheck(db database.DB, config config.Config) check {
	checks, err := health.New()
	if err != nil {
		panic(err)
	}

	return check{
		_db:     db,
		config:  config,
		_health: checks,
	}
}

func (h check) mustRegister(checker health.Config) {
	err := h._health.Register(checker)
	if err != nil {
		panic(err)
	}
}

func (h check) handler() http.HandlerFunc {
	return h.handlerFunc
}

// HandlerFunc is the HTTP handler function.
func (h check) handlerFunc(w http.ResponseWriter, r *http.Request) {
	c := h._health.Measure(r.Context())

	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(c)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(string(data)) //nolint:forbidigo
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	code := http.StatusOK
	if c.Status == health.StatusUnavailable {
		code = http.StatusServiceUnavailable
	}
	w.WriteHeader(code)
	_, _ = w.Write(data)
}
