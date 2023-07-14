package main

import (
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

func main() {
	getDeals()
}
