package utils

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func Success(w http.ResponseWriter, message string, data interface{}) {
	resp := APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	writeJSON(w, http.StatusOK, resp)
}

func Created(w http.ResponseWriter, message string, data interface{}) {
	resp := APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	writeJSON(w, http.StatusCreated, resp)
}

func Updated(w http.ResponseWriter, message string) {
	resp := APIResponse{
		Status:  "success",
		Message: message,
	}
	writeJSON(w, http.StatusOK, resp)
}

func Error(w http.ResponseWriter, message string, status int) {
	resp := APIResponse{
		Status: "error",
		Error:  message,
	}
	writeJSON(w, status, resp)
}

func writeJSON(w http.ResponseWriter, status int, resp APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}