package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Init mux router
	router := mux.NewRouter()

	// Route Handlers / Endpoints
	router.HandleFunc("/api/v1/user", addUser).Methods("POST")
	router.HandleFunc("/api/v1/users", getUsers).Methods("GET")
	router.HandleFunc("/api/v1/user", getUser).Methods("GET")

	router.HandleFunc("/api/v1/tastings", getTastings).Methods("GET")
	router.HandleFunc("/api/v1/tastings", addTasting).Methods("POST")

	println("Startup...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
