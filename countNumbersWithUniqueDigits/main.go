package main

import (
	"fmt"
)

func main() {
	fmt.Println(countNumbersWithUniqueDigits(2))
}

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1 // Only 0 satisfies the condition for n = 0
	}

	count := 10       // For n = 1, there are 10 unique digits (0 to 9)
	uniqueDigits := 9 // Start with 9 because the first digit cannot be 0

	for i := 2; i <= n; i++ {
		uniqueDigits *= 11 - i // Calculate the number of unique digits for each position
		count += uniqueDigits  // Add to the total count
	}

	return count
}
