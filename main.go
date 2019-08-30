package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Database
const (
	host     = "localhost"
	port     = 5432
	user     = "gobeeruser"
	password = "thisisthepassword"
	dbname   = "beerdb"
)

// Tasting Struct (Model)
type Tasting struct {
	ID    string `json:"id"`
	Theme string `json:"theme"`
	Beers []Beer `json:"beers"`
	Users []User `json:"users"`
	User  *User  `json:"user"`
}

// Beer struct
type Beer struct {
	ID        int    `json:"id"`
	Name      string `string:"name"`
	Tastingid int    `json:"tastingid"`
}

// Common user
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

// Scorecard for a user
type ScoreCard struct {
	ID       string   `json:"id"`
	Beers    []Beer   `json:"beers"`
	Comments []string `json:"comments"`
	Points   []int    `json:"points"`
}

func main() {
	// Init mux router
	router := mux.NewRouter()

	// Route Handlers / Endpoints
	router.HandleFunc("/api/v1/user", addUser).Methods("POST")
	router.HandleFunc("/api/v1/users", getUsers).Methods("GET")
	router.HandleFunc("/api/v1/user", getUser).Methods("GET")

	println("Startup...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

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

// Inserts new user into database and returns
func insertUserIntoDb(username string) User {
	insertUserStatement := "INSERT INTO users (username) VALUES ($1) RETURNING id"
	db := getDbConnection()
	defer db.Close()

	id := 0
	err := db.QueryRow(insertUserStatement, username).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println(id, username)
	var user User
	user.ID = id
	user.Username = username
	return user
}

func getAllUsersFromDb() []User {
	getUsersStatement := "SELECT * FROM users"
	db := getDbConnection()
	defer db.Close()

	rows, err := db.Query(getUsersStatement)
	if err != nil {
		panic(err)
	}
	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Username)
		if err != nil {
			panic(err)
		}
		fmt.Println(user)
		users = append(users, user)
	}
	return users
}

func getUserFromDb(id int) User {
	var username string
	getUserStatement := "SELECT * FROM users WHERE id=$1;"
	db := getDbConnection()
	defer db.Close()

	row := db.QueryRow(getUserStatement, id)
	switch err := row.Scan(&username, &id); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(id, username)
	default:
		panic(err)
	}
	var user User
	user.ID = id
	user.Username = username
	return user
}

func getDbConnection() *sql.DB {
	// Init database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
