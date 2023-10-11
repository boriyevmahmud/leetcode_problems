package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "15:04"
	fromTime := "09:00"
	toTime := "03:00"
	parsedFromTime, _ := time.Parse(layout, fromTime)
	parsedToTime, _ := time.Parse(layout, toTime)
	fmt.Println(parsedFromTime.After(parsedToTime))
	if parsedFromTime.After(parsedToTime) {
		parsedToTime = parsedToTime.Add(time.Hour * 24)
	}
	fmt.Println(parsedFromTime.After(parsedToTime))

	// currentDate := time.Now().Format("2006-01-02") // Get today's date in "YYYY-MM-DD" format

	// Combine the parsed times with today's date
	fromDateTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), parsedFromTime.Hour(), parsedFromTime.Minute(), 0, 0, time.Now().Location())
	toDateTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), parsedToTime.Hour(), parsedToTime.Minute(), 0, 0, time.Now().Location())

	fmt.Println(fromDateTime.Format("2006-01-02 15:04:05"))
	fmt.Println(toDateTime.Format("2006-01-02 15:04:05"))
}
