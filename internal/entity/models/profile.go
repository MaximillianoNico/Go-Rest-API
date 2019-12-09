package models

// ...Experience
type Experience struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	User     *User  `json:"user"`
	Position string `json:"position"`
	Division string `json:"division"`
}

// ...User
type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
