package main

import "fmt"

func main() {
	var numb string
	fmt.Scan(&numb)
	fmt.Println(romanToInt(numb))
}
func romanToInt(s string) int {
	var romanNumbers = make(map[string]int, 7)
	romanNumbers["I"] = 1
	romanNumbers["V"] = 5
	romanNumbers["X"] = 10
	romanNumbers["L"] = 50
	romanNumbers["C"] = 100
	romanNumbers["D"] = 500
	romanNumbers["M"] = 1000
	res := 0
	for i := 0; i <= len(s)-1; i++ {
		if i != len(s)-1 && romanNumbers[string(s[i])] < romanNumbers[string(s[i+1])] {
			res += romanNumbers[string(s[i+1])] - romanNumbers[string(s[i])]
			i++
			continue
		}
		res += romanNumbers[string(s[i])]
	}
	return res
}
