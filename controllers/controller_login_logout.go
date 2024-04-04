package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	m "tugas_explorasi_3_pbp/models"
)

// CheckUserLogin...
func CheckUserLogin(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	username := r.URL.Query()["username"]
	password := r.URL.Query()["password"]

	row := db.QueryRow("SELECT * FROM users WHERE username=? AND password=?", username[0], password[0])

	var user m.User
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		fmt.Println(err)
		sendResponse(w, http.StatusInternalServerError, "Login failed")
	} else {
		fmt.Println("Token will be generated")
		generateToken(w, r, user.ID, user.Username)
		sendResponse(w, http.StatusOK, "Login succeed")
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
