package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Create a new router using Gorilla Mux
	router := http.NewServeMux()

	// Define the handler function for the main URL for POST requests
	router.HandleFunc("/success", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("in here success")
		if r.Method == http.MethodPost {
			// Handle POST request
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error reading request body", http.StatusInternalServerError)
				return
			}
			defer r.Body.Close()

			// Process the request body for POST
			fmt.Printf("POST Request at Main URL: %s\n", r.URL.Path)
			fmt.Printf("Request Body: %s\n", body)

			// Respond with a message
			fmt.Fprintf(w, "POST Request at Main URL: %s\n", r.URL.Path)
		} else {
			// Handle other methods or respond with a "Method not allowed" error
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	// Define the handler function for the main URL for POST requests
	router.HandleFunc("/fail", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("in here fail")

		if r.Method == http.MethodPost {
			// Handle POST request
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error reading request body", http.StatusInternalServerError)
				return
			}
			defer r.Body.Close()

			// Process the request body for POST
			fmt.Printf("POST Request at Main URL1: %s\n", r.URL.Path)
			fmt.Printf("Request Body: %s\n", body)

			// Respond with a message
			fmt.Fprintf(w, "POST Request at Main URL1: %s\n", r.URL.Path)
		} else {
			// Handle other methods or respond with a "Method not allowed" error
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start the web server
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting the server: ", err)
	}
}
