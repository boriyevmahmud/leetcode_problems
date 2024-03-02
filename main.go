package main

import (
	"fmt"
	"net/http"
	"time"
)

func sendRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error sending request to %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Response from %s: %s\n", url, resp.Status)
}

func main() {
	url := "http://localhost:1237/v1/aggregate-customers"

	// Number of requests to send
	// numRequests := 2

	go sendRequest(url)
	time.Sleep(2 * time.Millisecond)
	sendRequest(url)
	time.Sleep(time.Second * 5)

	// Loop to send requests
	// for i := 0; i < numRequests; i++ {
	// 	sendRequest(url)
	// 	// time.Sleep(time.Second) // Wait for one second between requests
	// }
}
