package main

import (
	"fmt"
)

func main() {
	fmt.Println(titleToNumber("AB"))

}

func titleToNumber(columnTitle string) int {

	col := 0
	multi := 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		r := columnTitle[i]
		col += int(r-'A'+1) * multi
		multi *= 26
	}
	return col
}
