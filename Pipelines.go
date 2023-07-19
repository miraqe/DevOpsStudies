package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Config struct {
	PipedriveAPIToken string `json:"pipedrive_api_token"`
}

func loadConfig() (Config, error) {
	var config Config
	file, err := os.Open("config.json")
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return config, err
}

func forwardGetRequest(w http.ResponseWriter, r *http.Request) {
	var pipedriveAPIToken = os.Getenv("PIPEDRIVE_API_TOKEN")
	pipedriveURL := "https://api.pipedrive.com/v1/deals?api_token=" + pipedriveAPIToken

	// Create a new GET request to the Pipedrive API
	req, err := http.NewRequest(http.MethodGet, pipedriveURL, nil)
	if err != nil {
		log.Println("Error creating Pipedrive API request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

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

func addDealHandler(w http.ResponseWriter, r *http.Request) {
	var payloadData map[string]interface{}

	// Read the request body to get the payload data
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Unmarshal the request body into the payloadData map
	err = json.Unmarshal(body, &payloadData)
	if err != nil {
		log.Println("Error unmarshaling request body:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Call forwardAddDeal with the mock-up request, response, and payload data
	forwardAddDeal(w, r, payloadData)
}

func forwardAddDeal(w http.ResponseWriter, r *http.Request, payloadData map[string]interface{}) {
	//
	var pipedriveAPIToken = os.Getenv("PIPEDRIVE_API_TOKEN")
	pipedriveURL := "https://api.pipedrive.com/v1/deals?api_token=" + pipedriveAPIToken

	// Convert the payload data to JSON format
	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		log.Println("Error converting payload to JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Create a new POST request to the Pipedrive API with the payload
	req, err := http.NewRequest(http.MethodPost, pipedriveURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Println("Error creating a request to connect to Pipedrive API: ", err)
		http.Error(w, "Internal Server Error ", http.StatusInternalServerError)
		return
	}
	// Set the content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP client
	client := &http.Client{}

	// Send the request to the Pipedrive API
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request to Pipedrive API: ", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body from Pipedrive API
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body from Pipedrive API: ", err)
		http.Error(w, "Internal Server Error ", http.StatusInternalServerError)
		return

	}
	println("Connection Successful! Deal added: ", string(body))

	// Make sure response is in JSON and write the response to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func main() {
	http.HandleFunc("/myendpoint", forwardGetRequest)
	http.HandleFunc("/addDeal", addDealHandler)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
