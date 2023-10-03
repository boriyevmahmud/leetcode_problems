package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strStr("sadbutsad", "saad"))
}

func strStr(haystack string, needle string) int {

	return strings.Index(haystack, needle)
}
