package main

import (
	"net/http"
	"reflect"
	"testing"
)

func UnitTest(t *testing.T) {
	t.Run("Should be return Object", func(t *testing.T) {
		// Mock Data
		// experience = append(experience, Experience{
		// 	ID:    "124dfaw3aAg3",
		// 	Title: "Front End Engineer",
		// 	Desc:  "lorem ad gip sum",
		// 	User: &User{
		// 		Firstname: "Maximilliano",
		// 		Lastname:  "Lolong",
		// 		Email:     "maximilliano@okadoc.com",
		// 		Password:  "max1234",
		// 	},
		// 	Division: "IT Dev",
		// 	Position: "Junior Dev",
		// })
		var data []string
		expectData, _ := http.NewRequest("GET", "/v1/profiles", nil)

		if reflect.TypeOf(data) != reflect.TypeOf(expectData) {
			t.Fatalf("unexpected data type, should be return type data %v", reflect.TypeOf(data))
		}

	})
}
