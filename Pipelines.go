package main

import (
	"io"
	"log"
	"net/http"
)

func forwardGetRequest(w http.ResponseWriter, r *http.Request) {
	pipedriveURL := "https://api.pipedrive.com/v1/deals?api_token=a9bfb11e27d8e6f01b7d0e8d88a53cda91c454ac"

	// Create a new GET request to the Pipedrive API
	req, err := http.NewRequest(http.MethodGet, pipedriveURL, nil)
	if err != nil {
		log.Println("Error creating Pipedrive API request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Pipedrive API token as an Authorization header
	req.Header.Set("Authorization", "Bearer a9bfb11e27d8e6f01b7d0e8d88a53cda91c454ac")

	// Create a new HTTP client
	client := &http.Client{}

	// Send the request to the Pipedrive API
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request to Pipedrive API:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body from the Pipedrive API
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body from Pipedrive API:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	println("Connection Successful! Showing all deals: ", string(body))

	// Set the appropriate headers and write the response body to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func main() {
	http.HandleFunc("/myendpoint", forwardGetRequest)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
