package checks

import (
	"context"
	"time"

	"github.com/hellofresh/health-go/v4"
)

func (h check) db() health.Config {
	return health.Config{
		Name:      "DbConnection",
		Timeout:   2 * time.Second,
		SkipOnErr: false,
		Check: func(ctx context.Context) error {
			return h._db.Ping()
		},
	}
}
