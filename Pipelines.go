package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func getDeals() {
	// Set URL to connect to Pipedrive API with API token and parse the URL to format it properly
	endpointURL, err := url.Parse("https://api.pipedrive.com/api/v1/deals?api_token=a9bfb11e27d8e6f01b7d0e8d88a53cda91c454ac")
	method := "GET"
	if err != nil {
		fmt.Println("Error parsing URL: ", err)
		return
	}

	// Create HTTP client
	client := &http.Client{}

	// Create a new GET request to the endpoint URL
	req, err := http.NewRequest(method, endpointURL.String(), nil)
	if err != nil {
		fmt.Println("Error creating a new GET request to the endpoint URL!")
	}

	// Send the request and get a response
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error getting a response from the server: ", err)
		return
	}
	// Make sure response body is closed
	defer func() {
		err := res.Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error closing the response body: ", err)
	}

	// Print all deals that are found
	println("Connection Successful! Showing all deals: ", string(body))
}

func addDeal() {
	// Set URL to connect to Pipedrive API with API token and parse the URL to format it properly
	endpointUrl := "https://api.pipedrive.com/api/v1/deals?api_token=a9bfb11e27d8e6f01b7d0e8d88a53cda91c454ac"
	method := "POST"

	// Prepare the payload data for the new deal
	payloadData := map[string]interface{}{
		"title":              "Ingretchen",
		"value":              7209,
		"currency":           "EUR",
		"status":             "open",
		"org_id":             4,
		"participants_count": 1,
	}

	// Check if the title exists and is not empty
	title, titleExists := payloadData["title"]
	if !titleExists || title == "" {
		fmt.Println("Deal title is mandatory! You have not passed one. Try again by adding title name!")
		return
	}

	// Convert the payload data to JSON format
	payloadBytes, _ := json.Marshal(payloadData)

	// Create HTTP client
	client := &http.Client{}

	// Create a new POST request to the endpoint URL with the payload
	req, err := http.NewRequest(method, endpointUrl, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating a new GET request to the endpoint URL!")
	}
	// Set the content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Send the request and get a response
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	// Ensure the response body is closed after reading
	defer func() {
		err := res.Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}()
	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error closing the response body: ", err)
	}

	// Print the deal added
	println("Connection Successful! Deal added: ", string(body))
}

func updateDeal() {
	// Set URL to connect to Pipedrive API with API token and parse the URL to format it properly, determine deal value
	endpointUrl := "https://api.pipedrive.com/api/v1/deals/21?api_token=a9bfb11e27d8e6f01b7d0e8d88a53cda91c454ac"
	method := "PUT"

	// Change deal data
	payloadData := map[string]interface{}{
		"title":              "WinterBunches",
		"value":              7654,
		"currency":           "EUR",
		"status":             "open",
		"org_id":             4,
		"participants_count": 1,
	}

	// Convert the payload data to JSON format
	payloadBytes, _ := json.Marshal(payloadData)

	// Create HTTP client
	client := &http.Client{}

	// Create a new POST request to the endpoint URL with the payload
	req, err := http.NewRequest(method, endpointUrl, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating a new GET request to the endpoint URL!")
	}

	// Set the content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Send the request and get a response
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	// Ensure the response body is closed after reading
	defer func() {
		err := res.Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error closing the response body: ", err)
	}
	// Print changed deal
	println("Connection successful! Deal has been changed: ", string(body))
}

func main() {
	getDeals()
	addDeal()
	updateDeal()
}
