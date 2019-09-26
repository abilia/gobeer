package main

// Tasting :
type Tasting struct {
	ID    string `json:"id"`
	Theme string `json:"theme"`
	Beers []Beer `json:"beers"`
	Users []User `json:"users"`
	User  *User  `json:"user"`
}

// Beer :
type Beer struct {
	ID        int    `json:"id"`
	Name      string `string:"name"`
	Tastingid int    `json:"tastingid"`
}

// User : Common user
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

// ScoreCard : A user's scorecard
type ScoreCard struct {
	ID       string   `json:"id"`
	Beers    []Beer   `json:"beers"`
	Comments []string `json:"comments"`
	Points   []int    `json:"points"`
}
