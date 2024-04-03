package main

import (
	"database/sql"
	"fmt"
	"log"

	"tugas_explorasi_3_pbp/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jasonlvhit/gocron"
)

func queryDatabase() ([]string, error) {
	// Simulate database query
	// You would replace this with your actual database query logic

	// Open connection to the database
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_tugas_explorasi_3_pbp?parseTime=true&loc=Asia%2FJakarta")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Perform the database query
	rows, err := db.Query("SELECT username FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Process the query result
	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return names, nil
}

func main() {
	controllers.InitializeRedisClient()

	// Start HTTP server and handle login/connect routes
	controllers.Token()

	router := mux.NewRouter()

	router.HandleFunc("/login", controllers.CheckUserLogin).Methods("GET")

	s := gocron.NewScheduler()

	s.Every(1).Second().Do(func() {
		names, err := queryDatabase()
		if err != nil {
			log.Println("Error querying database:", err)
			return
		}
		fmt.Println("Query result:", names)
	})

	<-s.Start()
}
