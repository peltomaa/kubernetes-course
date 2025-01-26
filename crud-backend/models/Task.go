package models

import "time"

type Task struct {
	ID        int       `json:"id"`
	Task      string    `json:"task"`
	CreatedAt time.Time `json:"created_at"`
	IsDone    bool      `json:"is_done"`
}
