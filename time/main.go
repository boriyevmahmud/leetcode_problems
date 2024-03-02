package main

import (
	"fmt"
	"time"
)

func main() {
	fromDate := "2023-12-29 00:00:00"
	toDate := "2023-12-29 00:00:00"

	fromDateParsed, err := time.Parse("2006-01-02 15:04:05", fromDate)
	fmt.Println("error parsing fromDate: ", err)
	toDateParsed, err := time.Parse("2006-01-02 15:04:05", toDate)
	fmt.Println("error parsing toDate: ", err)

	// Create a range of dates
	var formattedDates []string
	for current := fromDateParsed; current.Before(toDateParsed) || current.Equal(toDateParsed); current = current.AddDate(0, 0, 1) {
		formattedDates = append(formattedDates, current.Format("2006-01-02"))
	}

	// Join the formatted dates with commas
	result := fmt.Sprintf("%s", formattedDates)
	fmt.Println(result)

}
