package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(ProcessPhone("+998900265560")) // на +998*****26560
	fmt.Println(ProcessPhone("+71112224455"))  // на +7*****224455
}

func ProcessPhone(phone string) string {
	var message string

	// Define a regular expression pattern to capture the country code
	countryCodeRegex := regexp.MustCompile(`^\+(\d{1,3})`)

	// Find the matching country code
	countryCodeMatches := countryCodeRegex.FindStringSubmatch(phone)
	if len(countryCodeMatches) < 2 {
		return "Invalid phone number"
	}
	countryCode := countryCodeMatches[1]

	// Apply logic based on the country code
	if len(phone) > len(countryCode)+4 { // Assuming the phone number has at least 4 more digits after the country code
		message = fmt.Sprintf(" на +%s*****%s", countryCode, phone[len(countryCode)+4:])
	}

	return message
}
