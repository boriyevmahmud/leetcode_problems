package main

import (
	"fmt"
	"internal/syscall/execenv"
)

func main() {
	fmt.Println(numberToTitle(1))

}

func numberToTitle(columnNumber int) string {

	col := ""
	for columnNumber > 0 {
		col = string(rune((columnNumber-1)%26+65)) + col
		columnNumber = (columnNumber - 1) / 26
	}
	
	return col
}
