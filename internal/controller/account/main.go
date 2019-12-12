package account

import (
	// db "./app/common/libs/db.go"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	models "github.com/MaximillianoNico/Go-Rest-API/internal/entity/models"
	"github.com/MaximillianoNico/Go-Rest-API/pkg/mongo"
)

var account []models.Account

func SignInWithEmailAndPass(w http.ResponseWriter, r *http.Request) {

	client, err := mongo.Connect()

	collection := client.Collection("Profile")

	if err != nil {
		log.Fatal(err)
	}

	newAccount := models.Account{
		Username: "ok",
		Password: "ok",
		Role:     "admin",
		Token:    "adawdawdawawd",
	}

	res, errDb := collection.InsertOne(context.TODO(), newAccount)

	if errDb != nil {
		log.Fatal(errDb)
	}
	fmt.Println("Inserted a Single Document: ", res.InsertedID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func SignInWithProvider(w http.ResponseWriter, r *http.Request) {
	//
}
