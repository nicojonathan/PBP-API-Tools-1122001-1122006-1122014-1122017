package models

import "time"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Task struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	StartTask time.Time `json:"start_task"`
	DueDate   time.Time `json:"due_date"`
	Details   string    `json:"details"`
	Notified  int      `json:"notified"`
	Email     string    `json:"email"`
}

type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type UsersResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

type TaskResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Task   `json:"data"`
}

type TasksResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Task `json:"data"`
}

type GeneralResponse struct {
	Status  int    `json:"status"`
	Message string `json:"messsage"`
}
