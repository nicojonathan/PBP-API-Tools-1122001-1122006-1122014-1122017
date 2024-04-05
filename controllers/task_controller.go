package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	m "tugas_explorasi_3_pbp/models"

	"github.com/go-redis/redis/v8"
)

func GetAllTasks() ([]m.Task, error) {
	db := connect()
	defer db.Close()

	// Perform the database query
	rows, err := db.Query("SELECT t.*, u.email FROM tasks t JOIN users u ON t.user_id=u.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Process the query result
	var tasks []m.Task
	for rows.Next() {
		var task m.Task
		if err := rows.Scan(&task.ID, &task.UserID, &task.Title, &task.StartTask, &task.DueDate, &task.Details, &task.Notified, &task.Email); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

    // Read from Request Body
    err := r.ParseForm()
    if err != nil {
        sendResponse(w, http.StatusBadRequest, "Error parsing form data")
        return
    }

    username := r.Form.Get("username")
    title := r.Form.Get("title")
    dueDate := r.Form.Get("due_date")
    details := r.Form.Get("details")

    ctx := context.Background()
    val, err := client.Get(ctx, username).Result()
    if err == redis.Nil {
		fmt.Println(err)
        sendResponse(w, http.StatusBadRequest, "Bad Request! You must log in first!")
        return
    } else if err != nil {
        sendResponse(w, http.StatusInternalServerError, "Internal Server Error")
        return
    }

    var values []string
    if err := json.Unmarshal([]byte(val), &values); err != nil {
        sendResponse(w, http.StatusInternalServerError, "Error unmarshaling JSON")
        return
    }

    userId := values[1]

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
    	sendResponse(w, http.StatusInternalServerError, "Failed to load time location")
	}

	// 2006-01-02 15:04:05 = layout untuk YYYY-MM-DD HH:mm:ss
    dueDateTimestamp, err := time.ParseInLocation("2006-01-02 15:04:05", dueDate, loc)
    if err != nil {
		fmt.Println(err)
        sendResponse(w, http.StatusBadRequest, "Error parsing due date")
        return
    }

    if title == "" || dueDate == "" || details == "" {
        sendResponse(w, http.StatusBadRequest, "Bad Request! You must include all parameters")
        return
    }

    startTask := time.Now()

    _, errQuery := db.Exec("INSERT INTO tasks (user_id, title, start_task, due_date, details, notified) VALUES (?,?,?,?,?,?)", 
        userId,
        title, 
        startTask, 
        dueDateTimestamp, 
        details, 
        0,
    )
    
    if errQuery != nil {
        log.Println(errQuery)
        sendResponse(w, http.StatusInternalServerError, "Invalid query")
        return
    }

    sendResponse(w, http.StatusOK, "Task added successfully")
}


func UpdateNotified(id int) {
	db := connect()
	defer db.Close()

	updateQuery := "UPDATE tasks SET notified = notified + 1 WHERE id = ?"
	db.Exec(updateQuery, id)
}