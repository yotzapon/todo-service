package entity

import (
	"time"
)

type TodoEntity struct {
	ID          string
	Title       string
	Description string
	IsCompleted bool
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
