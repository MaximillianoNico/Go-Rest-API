package profile

import (
	// db "./app/common/libs/db.go"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

var experience []Experience

func GetProfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(experience)
}

func GetProfileById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params

	for _, item := range experience {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Experience{})
}

func UpdateOrAddExperience(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var experiences Experience
	_ = json.NewDecoder(r.Body).Decode(&experiences)
	experiences.ID = strconv.Itoa(rand.Intn(10000000)) // mock data id - not safe
	experience = append(experience, experiences)

	json.NewEncoder(w).Encode(experience)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {

}

func UpdateEducation(w http.ResponseWriter, r *http.Request) {

}
