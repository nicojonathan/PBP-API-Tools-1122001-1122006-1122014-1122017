package controllers

import (
	"encoding/json"
	"net/http"

	m "tugas_explorasi_3_pbp/models"
)

func sendResponse(w http.ResponseWriter, status int, message string) {
	var response m.GeneralResponse
	response.Status = status
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
