package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	"tugas_explorasi_3_pbp/controllers"
	m "tugas_explorasi_3_pbp/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jasonlvhit/gocron"
)

func queryDatabase() ([]m.Task, error) {
	// Simulate database query
	// You would replace this with your actual database query logic

	// Open connection to the database
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_tugas_explorasi_3_pbp?parseTime=true&loc=Asia%2FJakarta")
	if err != nil {
		return nil, err
	}
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

func main() {
	// Initialize Redis client
	controllers.InitializeRedisClient()

	// 	// Start HTTP server and handle login/connect routes
	// 	//controllers.Token()

	// Start HTTP server in a separate goroutine
	go func() {
		router := mux.NewRouter()
		router.HandleFunc("/login", controllers.CheckUserLogin).Methods("GET")
		// router.HandleFunc("/tasks", ).Methods("GET")
		fmt.Println("Connected to port 8888")
		log.Println("Connected to port 8888")
		log.Fatal(http.ListenAndServe(":8888", router))
	}()

	// Start goCRON scheduler
	s := gocron.NewScheduler()
	s.Every(1).Second().Do(func() {

		result, err := queryDatabase()
		if err != nil {
			log.Println("Error querying database:", err)
			return
		}
		fmt.Println("Query result:", result)
		for _, task := range result {
			go func(t m.Task) {
				for {
					if time.Now().After(t.DueDate.Add(-10*time.Minute)) && t.Notified == 0 {
						fmt.Printf("Task " + t.Title + " is due in 10 minutes!")
						// Send email siniii
						// Query ke db update task notified++
						return
					} else if time.Now().After(t.DueDate.Add(-5*time.Minute)) && t.Notified == 1 {
						fmt.Printf("Task " + t.Title + " is due in 5 minutes!")
						// Send email siniii
						// Query ke db update task notified++
						return
					}
				}
			}(task)
		}

	})
	<-s.Start() // This line will block indefinitely, so it's typically not used in a real application
}

// func main() {
// 	controllers.InitializeRedisClient()

// 	// Start HTTP server and handle login/connect routes
// 	//controllers.Token()

// 	router := mux.NewRouter()

// 	router.HandleFunc("/login", controllers.CheckUserLogin).Methods("GET")

// 	http.Handle("/", router)

// 	fmt.Println("Connected to port 8888")
// 	log.Println("Connected to port 8888")
// 	log.Fatal(http.ListenAndServe(":8888", router))

// 	s := gocron.NewScheduler()

// 	s.Every(1).Second().Do(func() {
// 		names, err := queryDatabase()
// 		if err != nil {
// 			log.Println("Error querying database:", err)
// 			return
// 		}
// 		fmt.Println("Query result:", names)
// 	})

// 	<-s.Start()
// }
