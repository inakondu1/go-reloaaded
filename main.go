package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	text := "1E (hex) files were added"
	// in here i need to split hte word so i will be using strings.Fields. then i will range through the word. the if statement is me planning on replacing hex in the future run
	word := strings.Fields(text)
	for i, w := range word {
		if w == "(hex)" {
			previousword := word[i-1]
			// the lines of code below is where the conversion begine. this lines is expected to bring an out put of 30 after runing the code.
			decimal, err := strconv.ParseInt(previousword, 16, 64)
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Println(decimal)
		}
		// this 'i' meeans index, w is word. i want to print them out in an index formatt
		fmt.Println(i, w)
	}
	fmt.Println(word)

}
