package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestForwardGetRequest(t *testing.T) {
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
