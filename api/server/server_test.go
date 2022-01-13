package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	s := NewServer()

	// Testing the health check endpoint
	request, _ := http.NewRequest("GET", "/health", nil)
	response := httptest.NewRecorder()
	s.Router.ServeHTTP(response, request)

	// This request should always return status 200
	if http.StatusOK != response.Code {
		t.Errorf("Expected status code %d. Got %d", http.StatusOK, response.Code)
	}

	var jsonResult map[string]interface{}
	err := json.NewDecoder(response.Body).Decode(&jsonResult)
	if err != nil {
		t.Errorf("Unable to Decode the JSON into a map[string]interface{}")
	}
	// A valid response will always have a status:true in the body
	if jsonResult["status"] != true {
		t.Errorf("Expected result status %v. Got %v", true, jsonResult["status"])
	}
}

func TestCreateDeck(t *testing.T) {
	s := NewServer()

	// Testing the health check endpoint
	request, _ := http.NewRequest("POST", "/deck/create", nil)
	response := httptest.NewRecorder()
	s.Router.ServeHTTP(response, request)

	// This request should always return status 200
	if http.StatusOK != response.Code {
		t.Errorf("Expected status code %d. Got %d", http.StatusOK, response.Code)
	}

	var jsonResult map[string]interface{}
	err := json.NewDecoder(response.Body).Decode(&jsonResult)
	if err != nil {
		t.Errorf("Unable to Decode the JSON into a map[string]interface{}")
	}
	remaining := fmt.Sprintf("%v", jsonResult["remaining"])
	if err != nil {
		t.Errorf(err.Error())
	}

	if remaining != "52" {
		t.Errorf("Remaining cards after creation should be 52. Got %v", jsonResult["remaining"])
	}
}

// OMITTED OTHER TESTS TO AVOID A LOT OF CODE, BUT VERY SIMILAR TO THE TOP TEST
// NB: All tests are in each pkg as unit tests
