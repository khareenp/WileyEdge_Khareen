package main

import (
	"io"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}
//@localhost:8080/health-check
func main() {
	//Error func()
	http.HandleFunc("/health-check", HealthCheckHandler)
	http.ListenAndServe(":8080", nil)
}