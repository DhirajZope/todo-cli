package model

import "time"

type Task struct {
	ID          int
	ListID      int
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}
