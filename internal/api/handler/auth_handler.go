package handler

import "net/http"

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("auth handler ok"))
}
