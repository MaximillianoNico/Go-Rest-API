package main

import (
	"net/http"
	"reflect"
	"testing"
)

func UnitTest(t *testing.T) {
	t.Run("Should be return Object", func(t *testing.T) {
		var data []string
		expectData, _ := http.NewRequest("GET", "/v1/profiles", nil)

		if reflect.TypeOf(data) != reflect.TypeOf(expectData) {
			t.Fatalf("unexpected data type, should be return type data %v", reflect.TypeOf(data))
		}

	})
}
