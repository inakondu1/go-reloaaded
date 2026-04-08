package main

import (
	"fmt"
	"strings"
)
func main() {
	text := "1E (hex) files were added"
	word := strings.Fields(text)
	for i, w := range word {
// this i meeans index, w is word. i want to print them out in an index formatt
		fmt.Println(i, w)
	}
	   fmt.Println(word)

}