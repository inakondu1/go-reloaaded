package main

import (
	"fmt"
	"strconv"
)

func main() {
	hex := "1E"
	hexdecimal, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(hexdecimal)
}
