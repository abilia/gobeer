package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Add new user with username
func addUser(respWriter http.ResponseWriter, request *http.Request) {
	respWriter.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(request.Body).Decode(&user)
	insertedUser := insertUserIntoDb(user.Username)
	json.NewEncoder(respWriter).Encode(insertedUser)
}

// Return all current users
func getUsers(respWriter http.ResponseWriter, request *http.Request) {
	respWriter.Header().Set("Content-Type", "application/json")
	users := getAllUsersFromDb()
	json.NewEncoder(respWriter).Encode(users)
}

// Get user with specific id
func getUser(respWriter http.ResponseWriter, request *http.Request) {
	respWriter.Header().Set("Content-Type", "application/json")
	keys, ok := request.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		return
	}
	id, err := strconv.Atoi(keys[0])
	if err != nil {
		panic(err)
	}
	user := getUserFromDb(id)
	json.NewEncoder(respWriter).Encode(user)
}

// Get all tastings :
func getTastings(respWriter http.ResponseWriter, request *http.Request) {
	respWriter.Header().Set("Content-Type", "application/json")
	response := getAllTastings()
	json.NewEncoder(respWriter).Encode(response)
}

func getTasting(respWriter http.ResponseWriter, request *http.Request) {
	respWriter.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	response := getTastingByID(id)
	json.NewEncoder(respWriter).Encode(response)
}

func addTasting(respWriter http.ResponseWriter, request *http.Request) {
	respWriter.Header().Set("Content-Type", "application/json")
	var tasting Tasting
	_ = json.NewDecoder(request.Body).Decode(&tasting)
	insertTasting(tasting.Name)
}
