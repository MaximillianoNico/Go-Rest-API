package profile

import (
	// db "./app/common/libs/db.go"
	"encoding/json"
	models "github.com/MaximillianoNico/Go-Rest-API/internal/entity/models"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

var experience []models.Experience

func GetProfiles(w http.ResponseWriter, r *http.Request) {
	experience := append(experience, models.Experience{
		ID:    "124dfaw3aAg3",
		Title: "Front End Engineer",
		Desc:  "lorem ad gip sum",
		User: &models.User{
			Firstname: "Maximilliano",
			Lastname:  "Lolong",
			Email:     "maximilliano@okadoc.com",
			Password:  "max1234",
		},
		Division: "IT Dev",
		Position: "Junior Dev",
	})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(experience)
}

func GetProfileById(w http.ResponseWriter, r *http.Request) {
	experience := append(experience, models.Experience{
		ID:    "124dfaw3aAg3",
		Title: "Front End Engineer",
		Desc:  "lorem ad gip sum",
		User: &models.User{
			Firstname: "Maximilliano",
			Lastname:  "Lolong",
			Email:     "maximilliano@okadoc.com",
			Password:  "max1234",
		},
		Division: "IT Dev",
		Position: "Junior Dev",
	})

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params

	for _, item := range experience {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&models.Experience{})
}

func UpdateOrAddExperience(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var experiences models.Experience
	_ = json.NewDecoder(r.Body).Decode(&experiences)
	experiences.ID = strconv.Itoa(rand.Intn(10000000)) // mock data id - not safe
	experience = append(experience, experiences)

	json.NewEncoder(w).Encode(experience)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {

}

func UpdateEducation(w http.ResponseWriter, r *http.Request) {

}
