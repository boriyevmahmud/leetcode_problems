package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiKey = "6987bf6f-2a3a-42eb-9de4-b79bd7381b5e"

type YandexDirectionsResponse struct {
	Routes []struct {
		Distance struct {
			Value int `json:"value"`
		} `json:"distance"`
	} `json:"routes"`
}

func main() {
	// Example latitude and longitude points (replace with your coordinates)
	origin := "69.2444957354736, 69.2444957354736"      //Ташкент, улица Юсуфа Хамадони, 92
	destination := "69.238962,41.269604" // St. Petersburg, Russia

	url := fmt.Sprintf("https://api-maps.yandex.ru/services/route/2.0/?apikey=%s&lang=en_US&origin=%s&destination=%s", apiKey, origin, destination)

	// Send the request to the Yandex Maps API
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Parse the JSON response
	var directionsResponse YandexDirectionsResponse
	if err := json.Unmarshal(body, &directionsResponse); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Extract the road distance
	if len(directionsResponse.Routes) > 0 {
		distanceValue := directionsResponse.Routes[0].Distance.Value
		distanceKm := float64(distanceValue) / 1000.0

		fmt.Printf("Road distance: %.2f km\n", distanceKm)
	} else {
		fmt.Println("No route found.")
	}
}
