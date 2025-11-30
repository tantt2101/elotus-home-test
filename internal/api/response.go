package api

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
