package main

import (
	"fmt"
	"time"
)

func main() {
	var f int32 = 10
	layout := "15:04"
	parsedTime, _ := time.Parse(layout, "15:54")
	currentTime := time.Now()
	currentDate := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	resultTime := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), parsedTime.Hour(), parsedTime.Minute(), 0, 0, currentDate.Location())
	resultTime = resultTime.Add(-time.Minute * time.Duration(f))


	// Compare the resulting time with the current time
	check := currentTime.After(resultTime)

	fmt.Println("resultTime:", resultTime)
	fmt.Println("check:", check)
}
