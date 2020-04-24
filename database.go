package main

import (
	"database/sql"
	"fmt"
)

// Database
const (
	host     = "localhost"
	port     = 5454
	user     = "gobeeruser"
	password = "thisisthepassword"
	dbname   = "beerdb"
)

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

func getAllTastings() []Tasting {
	db := getDbConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tastings")
	if err != nil {
		panic(err)
	}

	var tastings []Tasting = make([]Tasting, 0)
	for rows.Next() {
		var tasting Tasting
		rows.Scan(&tasting.ID, &tasting.Name)
		tastings = append(tastings, tasting)
	}
	return tastings
}

func getTastingByID(id int) (Tasting, error) {
	res := Tasting{}
	var name string

	db := getDbConnection()
	defer db.Close()

	err := db.QueryRow("SELECT * FROM tastings WHERE id = $1", id).Scan(&id, &name)
	if err == nil {
		res = Tasting{ID: id, Name: name}
	}
	return res, err
}

func insertTasting(name string) int {
	db := getDbConnection()
	defer db.Close()

	id := 0
	err := db.QueryRow("INSERT INTO tastings (name) VALUES ($1) RETURNING id", name).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}

func deleteTastingByID(id int) {
	db := getDbConnection()
	defer db.Close()

	_, err := db.Exec("DELETE FROM tastings WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
}

func getBeersByTastingID(tastingID int) []Beer {
	db := getDbConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM beers WHERE tastingID = $1", tastingID)
	if err != nil {
		panic(err)
	}

	var beers []Beer = make([]Beer, 0)
	for rows.Next() {
		var beer Beer
		rows.Scan(&beer.ID, &beer.Name, &beer.TastingID)
		beers = append(beers, beer)
	}
	return beers
}

func insertBeer(tastingID int, beerName string) {
	db := getDbConnection()
	defer db.Close()

	id := 0
	err := db.QueryRow("INSERT INTO beers (name, tastingID) VALUES ($1, $2) RETURNING id", beerName, tastingID).Scan(&id)
	if err != nil {
		panic(err)
	}
}
