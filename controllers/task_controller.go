package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	m "tugas_explorasi_3_pbp/models"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM users"

	// Read from Query Param
	username := r.URL.Query()["username"]

	if username != nil {
		fmt.Println(username[0])
		query += " WHERE username='" + username[0] + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		// send error response
		sendResponse(w, http.StatusInternalServerError, "Invalid query")
		return
	}

	var user m.User
	var users []m.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			log.Println(err)
			// send error response
			sendResponse(w, http.StatusInternalServerError, "Failed to scan data")
			return
		} else {
			users = append(users, user)
		}
	}

	var response m.UsersResponse
	if len(users) > 0 {
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Data = users
	} else {
		response.Status = http.StatusNotFound
		response.Message = "No users found"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
