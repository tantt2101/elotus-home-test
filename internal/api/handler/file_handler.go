package handler

import "net/http"

func FileHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("auth handler ok"))
}
