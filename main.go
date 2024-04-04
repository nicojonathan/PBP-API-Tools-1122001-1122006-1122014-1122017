package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strconv"
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

func sendMail(email string, subject string, message string, due time.Time) {
	auth := smtp.PlainAuth(
		"",
		"gantengnic16@gmail.com", //ini tuh admin (pengirim)
		"eshlwrkasazelgnz",
		"smtp.gmail.com",
	)

	msg := fmt.Sprintf("Subject: %s\n\n%s\n%s", subject, message, "This task is due in " + strconv.Itoa(int(time.Until(due).Minutes())) + "  minutes")


	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"gantengnic16@gmail.com",
		[]string{email}, //penerima notifikasi
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
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
	s.Every(10).Second().Do(func() {

		result, err := queryDatabase()
		if err != nil {
			log.Println("Error querying database:", err)
			return
		}
		fmt.Println("Query result:", result)
		for _, task := range result {
			go func(t m.Task) {
					if time.Now().After(t.DueDate.Add(-10*time.Minute)) && time.Now().Before(t.DueDate) && t.Notified == 0 {
						//fmt.Printf("Task " + t.Title + " is due in 10 minutes!")

						sendMail(t.Email, t.Title, t.Details, t.DueDate)
						controllers.UpdateNotified(t.ID)

						// Query ke db update task notified++
						return
					} else if time.Now().After(t.DueDate.Add(-5*time.Minute)) && time.Now().Before(t.DueDate) && t.Notified == 1 {
						//fmt.Printf("Task " + t.Title + " is due in 5 minutes!")

						sendMail(t.Email, t.Title, t.Details, t.DueDate)
						controllers.UpdateNotified(t.ID)

						return
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
