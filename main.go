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
	var day = 0
	fmt.Println(time.Weekday(day))
}
