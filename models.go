package main

// Tasting :
type Tasting struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Beer :
type Beer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	TastingID int    `json:"tastingId"`
}

// Picture :
type Picture struct {
	UUID string `json:"id"`
}

// User : Common user
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

// ScoreCard : A user's scorecard
type ScoreCard struct {
	UserID User `json:"userId"`
	Beers  Beer `json:"beerId"`
	Points int  `json:"points"`
}
