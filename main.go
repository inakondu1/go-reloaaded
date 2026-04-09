package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Examples of numbers in different systems
	hexNum := "1E"
	binNum := "11110"
	decNum := "30"

	// Convert hex to decimal
	hexDec, _ := strconv.ParseInt(hexNum, 16, 64)
	fmt.Println("Hex", hexNum, "-> Decimal", hexDec)

	// Convert binary to decimal
	binDec, _ := strconv.ParseInt(binNum, 2, 64)
	fmt.Println("Binary", binNum, "-> Decimal", binDec)

	// Convert decimal string to integer
	dec, _ := strconv.ParseInt(decNum, 10, 64)
	fmt.Println("Decimal", decNum, "-> Decimal", dec)
}
