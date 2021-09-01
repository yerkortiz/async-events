package server

import (
	"net/http"

	m2 "milestone2/handlers"

	"github.com/gorilla/mux"
)

func RunServer() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", m2.HealthCheck).Methods("GET")
	http.ListenAndServe(":8080", r)
}
