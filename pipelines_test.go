package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func apiToken(t *testing.T) {
	config, err := loadConfig()
	if err != nil {
		t.Fatalf("Error loading config: %s", err)
	}

	// Set the API token as an environment variable
	os.Setenv("PIPEDRIVE_API_TOKEN", config.PipedriveAPIToken)
}
func TestForwardGetRequest(t *testing.T) {
	apiToken(t)

	// Create a mock-up http.ResponseWriter
	w := httptest.NewRecorder()

	// Create a mock-up http.Request for the GET request
	r := httptest.NewRequest(http.MethodGet, "/myendpoint", nil)

	// Call forwardGetRequest with the mock-up request and response
	forwardGetRequest(w, r)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	// Check if the response body contains the "title" field
	expectedTitle := "Ingretchen"
	if !strings.Contains(w.Body.String(), expectedTitle) {
		t.Errorf("Expected title '%s' not found in the response", expectedTitle)
	}
}

func TestForwardAddDeal(t *testing.T) {
	apiToken(t)

	// Prepare the payload data for the new deal
	payloadData := map[string]interface{}{
		"title":              "LongerDivine",
		"value":              2775,
		"currency":           "EUR",
		"status":             "open",
		"org_id":             2,
		"participants_count": 1,
	}

	// Convert the payload data to JSON format
	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		t.Fatalf("Error converting payload to JSON: %s", err)
	}

	// Create a new POST request with the payload
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/addDeal", bytes.NewBuffer(payloadBytes))
	r.Header.Set("Content-Type", "application/json")

	// Call forwardAddDeal with the mock-up request, response, and payload data
	forwardAddDeal(w, r, payloadData)

	// Check the response status code
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, w.Code)
	}
}
