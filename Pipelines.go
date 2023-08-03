package main

import (
	"bytes"
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// Collect metrics for monitoring and observability purposes.
func collectMetrics() {
	go func() {
		for {
			opsProcessed.Inc() // Increment the total number of HTTP requests
			time.Sleep(2 * time.Second)
			duration := rand.Float64()
			opsDuration.Observe(duration) // Record the duration of HTTP requests
			time.Sleep(time.Second)
		}
	}()
}

// Create Prometheus custom metrics to collect and expose for monitoring.
var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "httpTotalRequests",
		Help: "TOTAL NUMBER OF HTTP REQUESTS",
	})
)

var (
	opsDuration = promauto.NewSummary(prometheus.SummaryOpts{
		Name: "httpDuration",
		Help: "HTTP REQUEST DURATION IN SECONDS",
	})
)

// Add a function Config, which represents configuration from a JSON file.
type Config struct {
	PipedriveAPIToken string `json:"PipedriveAPIToken"`
}

// Add a function loadConfig, which loads the configuration key from the "config.json" file.
func loadConfig() (Config, error) {
	var config Config

	// Check if the PIPEDRIVE_API_TOKEN environment variable is set
	apiToken := os.Getenv("PIPEDRIVE_API_TOKEN")
	if apiToken != "" {
		log.Println("Using API Token from environment variable:", apiToken)
		config.PipedriveAPIToken = apiToken
		return config, nil
	}

	// Read from "config.json" if the environment variable is not set
	file, err := os.Open("config.json")
	if err != nil {
		log.Println("Error opening config.json:", err)
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Println("Error decoding config.json:", err)
	}
	log.Println("Using API Token from config.json:", config.PipedriveAPIToken)
	return config, err
}

// getDealsHandler handles the HTTP request for getting deals from the Pipedrive API.
func getDealsHandler(w http.ResponseWriter, r *http.Request) {

	// Load the configuration from config.json
	config, err := loadConfig()
	if err != nil {
		log.Println("Error loading config:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	pipedriveURL := "https://api.pipedrive.com/v1/deals?api_token=" + config.PipedriveAPIToken

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

	// Set the appropriate headers and write the response body to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

// addDealHandler handles the HTTP request for adding a new deal to the Pipedrive API.
func addDealHandler(w http.ResponseWriter, r *http.Request) {

	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check if the Content-Type header is set to "application/json"
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Invalid Content-Type. Expected application/json", http.StatusUnsupportedMediaType)
		return
	}

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
		log.Println("Error un-marshaling request body:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Call addDeal with the mock-up request, response, and payload data
	addDeal(w, r, payloadData)
}

// addDeal adds a new deal to the Pipedrive API using the provided payload data.
func addDeal(w http.ResponseWriter, r *http.Request, payloadData map[string]interface{}) {
	// Load the configuration from config.json
	config, err := loadConfig()
	if err != nil {
		log.Println("Error loading config:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	pipedriveURL := "https://api.pipedrive.com/v1/deals?api_token=" + config.PipedriveAPIToken

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
		log.Println("Error reading response from Pipedrive API: ", err)
		http.Error(w, "Internal Server Error ", http.StatusInternalServerError)
		return
	}
	println("Connection Successful! Deal added: ", string(body))

	// Make sure response is in JSON and write the response to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

// changeDealHandler handles the HTTP request for changing an existing deal in the Pipedrive API.
func changeDealHandler(w http.ResponseWriter, r *http.Request) {

	var payloadData map[string]interface{}

	// Read the request body to get the payload data
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading the request body: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &payloadData)
	if err != nil {
		log.Println("Error un-marshaling request body: ", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	changeDeal(w, r, payloadData)
}

// changeDeal changes an existing deal in the Pipedrive API using the provided payload data.
func changeDeal(w http.ResponseWriter, r *http.Request, payloadData map[string]interface{}) {
	// Load the configuration from config.json
	config, err := loadConfig()
	if err != nil {
		log.Println("Error loading config:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	pipedriveURL := "https://api.pipedrive.com/v1/deals/130?api_token=" + config.PipedriveAPIToken

	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		log.Println("Error marshaling Payload Data: ", err)
		http.Error(w, "Internal Server Error: ", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest(http.MethodPut, pipedriveURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Println("Error creating a new request to Pipedrive API: ", err)
		http.Error(w, "Internal Server Error: ", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error reading a response from Pipedrive API: ", err)
		http.Error(w, "Internal Server Error: ", http.StatusInternalServerError)
		return
	}
	// Read the response body from Pipedrive API
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response from Pipedrive API: ", err)
		http.Error(w, "Internal Server Error ", http.StatusInternalServerError)
		return
	}
	println("Connection Successful! Deal changed: ", string(body))

	// Make sure response is in JSON and write the response to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func main() {

	collectMetrics()

	http.HandleFunc("/getDeals", getDealsHandler)
	http.HandleFunc("/addDeal", addDealHandler)
	http.HandleFunc("/changeDeal", changeDealHandler)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Server started on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
