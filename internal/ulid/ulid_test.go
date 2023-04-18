package ulid_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yotzapon/todo-service/internal/ulid"
)

func TestNew(t *testing.T) {
	t.Run("return a uuid with prefix", func(t *testing.T) {
		result := ulid.New("prefix")
		assert.True(t, strings.HasPrefix(result, "prefix"))
	})
}
