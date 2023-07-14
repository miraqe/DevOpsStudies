package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func getDeals() {
	baseURL := "https://api.pipedrive.com/api/v1/deals?api_token=a9bfb11e27d8e6f01b7d0e8d88a53cda91c454ac"
	method := "GET"

	endpointUrl, _ := url.Parse(baseURL)

	client := &http.Client{}
	req, _ := http.NewRequest(method, endpointUrl.String(), nil)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}()

	body, _ := io.ReadAll(res.Body)
	println("Connection Successful! Showing all deals: ", string(body))
}

func main() {
	getDeals()
}
