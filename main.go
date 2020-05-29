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
	router.HandleFunc("/api/v1/tastings/{id:[0-9]+}", getTasting).Methods("GET")
	router.HandleFunc("/api/v1/tastings", addTasting).Methods("POST")
	router.HandleFunc("/api/v1/tastings/{id:[0-9]+}", deleteTasting).Methods("DELETE")

	router.HandleFunc("/api/v1/tastings/{tastingId:[0-9]+}/beers", getBeers).Methods("GET")
	router.HandleFunc("/api/v1/tastings/{tastingId:[0-9]+}/beers", addBeer).Methods("POST")

	router.HandleFunc("/api/v1/coronabeers", getAllCoronaBeers).Methods("GET")
	router.HandleFunc("/api/v1/coronabeers/{id:[0-9]+}", getCoronaBeer).Methods("GET")
	router.HandleFunc("/api/v1/coronabeers/{id:[0-9]+}", editCoronaBeer).Methods("PUT")
	router.HandleFunc("/api/v1/coronabeers", addCoronaBeer).Methods("POST")

	router.HandleFunc("/api/v1/pictures", uploadFile).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	println("Startup...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
