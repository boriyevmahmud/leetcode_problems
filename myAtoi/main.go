package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	s := "w10734368512"
	fmt.Println(myAtoi(s))
}
func myAtoi(s string) int {
	numbers := make(map[string]bool)
	st := make(map[string]bool)
	numbers["0"] = true
	numbers["1"] = true
	numbers["2"] = true
	numbers["3"] = true
	numbers["4"] = true
	numbers["5"] = true
	numbers["6"] = true
	numbers["7"] = true
	numbers["8"] = true
	numbers["9"] = true
	st["-"] = true
	st["+"] = true
	str := ""
	if s == ""{
		return 0
	}
	if !numbers[string(s[0])] && !st[string(s[0])] && (string(s[0]) !=" " || string(s[0])=="."){
		return 0
	}
	for i := 0; i < len(s); i++ {
		fmt.Println("string(s[i])", string(s[i]))
		if i >= 1 && str != "" && string(s[i]) == " " {
			fmt.Println("in here")
			break
		}
		if string(s[i]) == " " {
			fmt.Println("in 2")

			continue
		}
		

		if i >= 1 && !st[string(s[i])] && !numbers[string(s[i])] {
			fmt.Println("in 4")

			break
		}
		if i >= 1 && st[string(s[i])] && numbers[string(s[i-1])] {
			fmt.Println("in 5")
			break
		}
		if numbers[string(s[i])] || st[string(s[i])] {
			fmt.Println("in 6")

			str += string(s[i])
		}
	}
	fmt.Println("str", str)
	res, _ := strconv.Atoi(str)

	if res > math.MaxInt32 {
		return math.MaxInt32
	} else if res < -math.MaxInt32-1 {
		return -math.MaxInt32 - 1
	}

	return res
}
