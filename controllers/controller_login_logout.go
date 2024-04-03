package controllers

import (
	"encoding/json"
	"net/http"

	m "tugas_explorasi_3_pbp/models"
)

// CheckUserLogin...
func CheckUserLogin(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	name := r.URL.Query()["name"]

	row := db.QueryRow("SELECT * FROM users WHERE name=?", name[0])

	var user m.User
	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		sendResponse(w, http.StatusInternalServerError, "Login failed")
	} else {
		generateToken(w, user.ID, user.Username)
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
