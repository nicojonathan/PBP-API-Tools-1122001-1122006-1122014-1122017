package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	m "tugas_explorasi_3_pbp/models"

	"github.com/go-redis/redis/v8"
)

// CheckUserLogin...
func CheckUserLogin(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	username := r.URL.Query()["username"]
	password := r.URL.Query()["password"]

	ctx := context.Background()
	val, err := client.Get(ctx, username[0]).Result()
	if val == "" || err == redis.Nil {
		row := db.QueryRow("SELECT * FROM users WHERE username=? AND password=?", username[0], password[0])

		var user m.User
		if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
			fmt.Println(err)
			if err == sql.ErrNoRows {
				sendResponse(w, http.StatusBadRequest, "Invalid username / password!")
				return
			}
			sendResponse(w, http.StatusInternalServerError, "Login failed")
		} else {
			fmt.Println("Token will be generated")
			generateToken(w, user.ID, user.Username)
			sendResponse(w, http.StatusOK, "Login succeed")
		}
	} else {
		sendResponse(w, http.StatusBadRequest, "You have already logged in!")
	}

}

// Logout...
func Logout(w http.ResponseWriter, r *http.Request) {
	resetUserToken(w)

	var response m.UserResponse
	response.Status = 200
	response.Message = "Success"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
