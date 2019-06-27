package main

import (
	// db "./app/common/libs/db.go"
	"encoding/json"
	"math/rand"
	"log"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
)

// ...Experience
type Experience struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	User   	*User 	`json:"user"`
	Position string `json:"position"`
	Division string `json:"division"`
}

// ...User
type User struct {
	Firstname string `json:"firstname"`
	Lastname string	`json:"lastname"`
	Email string `json:"email"`
	Password string `json:"password"`
}

var experience []Experience

func getProfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(experience)
}

func getProfileById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r); // get params

	for _, item := range experience {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Experience{})
}

func updateOrAddExperience(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var experiences Experience
	_ = json.NewDecoder(r.Body).Decode(&experiences)
	experiences.ID = strconv.Itoa(rand.Intn(10000000)) // mock data id - not safe
	experience = append(experience, experiences)
	
	json.NewEncoder(w).Encode(experience)
}

func updateProject(w http.ResponseWriter, r *http.Request) {

}

func updateEducation(w http.ResponseWriter, r *http.Request) {

}

func main() {

	// db.connection()
	r := mux.NewRouter()

	// Mock Data
	experience = append(experience, Experience{ 
		ID: "124dfaw3aAg3", 
		Title: "Front End Engineer", 
		Desc: "lorem ad gip sum",
		User: &User{
			Firstname: "Maximilliano",
			Lastname: "Lolong",
			Email: "maximilliano@okadoc.com",
			Password: "max1234",
		},
		Division: "IT Dev",
		Position: "Junior Dev",
	})

	experience = append(experience, Experience{ 
		ID: "124dfaw1SfAg3", 
		Title: "Back End Engineer", 
		Desc: "lorem ad gip sum",
		User: &User{
			Firstname: "Maximilliano",
			Lastname: "Lolong",
			Email: "maximilliano@okadoc.com",
			Password: "max1234",
		},
		Division: "IT Dev",
		Position: "Junior Dev",
	})

	r.HandleFunc("/v1/profiles", getProfiles).Methods("GET")
	r.HandleFunc("/v1/profile/{id}", getProfileById).Methods("GET")
	r.HandleFunc("/v1/profile/experience", updateOrAddExperience).Methods("POST")
	r.HandleFunc("/v1/profile/project", updateProject).Methods("POST")
	r.HandleFunc("/v1/profile/education", updateEducation).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}