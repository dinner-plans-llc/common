package server

import "net/http"

// Healthz
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{ alive: true }"))
}
