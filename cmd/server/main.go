package main

import (
	"log"
	"net/http"

	// profile "github.com/MaximillianoNico/Go-Rest-API/modules/profile"
	// "github.com/MaximillianoNico/Go-Rest-API/pkg/mongo"
	"github.com/MaximillianoNico/Go-Rest-API/internal/controller/profile"
	"github.com/MaximillianoNico/Go-Rest-API/internal/controller/account"
	"github.com/MaximillianoNico/Go-Rest-API/internal/config/env"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	
	// Routes Profile User
	r.HandleFunc("/v1/profiles", profile.GetProfiles).Methods("GET")
	r.HandleFunc("/v1/profile/{id}", profile.GetProfileById).Methods("GET")
	r.HandleFunc("/v1/profile/experience", profile.UpdateOrAddExperience).Methods("POST")
	r.HandleFunc("/v1/profile/project", profile.UpdateProject).Methods("POST")
	r.HandleFunc("/v1/profile/education", profile.UpdateEducation).Methods("POST")

	// Routes Authentication
	r.HandleFunc("/v1/auth/signin", account.SignInWithEmailAndPass).Methods("POST")

	log.Fatal(http.ListenAndServe(env.Get(env.HTTPListenPort), r))
}
