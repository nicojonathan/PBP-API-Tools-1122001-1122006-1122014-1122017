package main

import (
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

func sendMail(email string, subject string, message string, due time.Time) {
	auth := smtp.PlainAuth(
		"",
		"gantengnic16@gmail.com", //ini tuh admin (pengirim)
		"eshlwrkasazelgnz",
		"smtp.gmail.com",
	)

	msg := fmt.Sprintf("Subject: %s\n\n%s\n%s", "REMINDER!!! " + subject, message, "This task is due in " + strconv.Itoa(int(time.Until(due).Minutes())) + "  minutes")


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

	// Start HTTP server in a separate goroutine
	go func() {
		router := mux.NewRouter()

		router.HandleFunc("/tasks", controllers.AddTask).Methods("POST")
		router.HandleFunc("/login", controllers.CheckUserLogin).Methods("GET")

		fmt.Println("Connected to port 8888")
		log.Println("Connected to port 8888")
		log.Fatal(http.ListenAndServe(":8888", router))
	}()

	// Start goCRON scheduler
	s := gocron.NewScheduler()
	s.Every(10).Second().Do(func() {

		result, err := controllers.GetAllTasks()
		if err != nil {
			log.Println("Error querying database:", err)
			return
		}

		fmt.Println("QUERY RESULT")
		for _, task := range result {
			fmt.Printf("Task ID: %d\n", task.ID)
			fmt.Printf("User ID: %d\n", task.UserID)
			fmt.Printf("Title: %s\n", task.Title)
			fmt.Printf("Start Task: %s\n", task.StartTask.Format("2006-01-02 15:04:05"))
			fmt.Printf("Due Date: %s\n", task.DueDate.Format("2006-01-02 15:04:05"))
			fmt.Printf("Details: %s\n", task.Details)
			fmt.Printf("Notified: %d\n", task.Notified)
			fmt.Printf("Email: %s\n", task.Email)
			fmt.Println("-------------------------------------")

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
	<-s.Start()
}
