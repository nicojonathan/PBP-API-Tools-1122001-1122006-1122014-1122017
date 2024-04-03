package models

import "time"

type Task struct {
	ID       int
	Title    string
	DueDate  time.Time
	Email    string
	Notified bool
}
