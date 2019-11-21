package auth

// db "./app/common/libs/db.go"

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
