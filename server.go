package main

import (
	"log"
	"net/http"

	profile "github.com/MaximillianoNico/GO-REST-API/modules/profile"

	"github.com/gorilla/mux"
)

func main() {

	// db.connection()
	r := mux.NewRouter()

	// Routes Profile User
	r.HandleFunc("/v1/profiles", profile.GetProfiles).Methods("GET")
	r.HandleFunc("/v1/profile/{id}", profile.GetProfileById).Methods("GET")
	r.HandleFunc("/v1/profile/experience", profile.UpdateOrAddExperience).Methods("POST")
	r.HandleFunc("/v1/profile/project", profile.UpdateProject).Methods("POST")
	r.HandleFunc("/v1/profile/education", profile.UpdateEducation).Methods("POST")

	// Routes Authentication

	log.Fatal(http.ListenAndServe(":8080", r))
}
