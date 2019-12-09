package account

import (
	// db "./app/common/libs/db.go"
	"encoding/json"
	"net/http"

	models "github.com/MaximillianoNico/Go-Rest-API/internal/entity/models"
)

var account []models.Account

func SignInWithEmailAndPass(w http.ResponseWriter, r *http.Request) {
	login := append(account, models.Account{
		Username: "ok",
		Password: "ok",
		Role:     "admin",
		Token:    "adawdawdawawd",
	})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(login)
}

func SignInWithProvider(w http.ResponseWriter, r *http.Request) {
	//
}
