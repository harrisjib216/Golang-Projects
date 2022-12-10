package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Odd or Even
//
//  1. Get a number from the command-line.
//
//  2. Find whether the number is odd, even and divisible by 8.
//
// RESTRICTION
//  Handle the error. If the number is not a valid number,
//  or it's not provided, let the user know.
//
// EXPECTED OUTPUT
//  go run main.go 16
//    16 is an even number and it's divisible by 8
//
//  go run main.go 4
//    4 is an even number
//
//  go run main.go 3
//    3 is an odd number
//
//  go run main.go
//    Pick a number
//
//  go run main.go ABC
//    "ABC" is not a number
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Pick a number")
	} else if num, err := strconv.Atoi(os.Args[1]); err != nil {
		fmt.Printf("%q is not a number", os.Args[1])
	} else if num%2 == 0 {
		fmt.Printf("%d is an even number", num)

		if num%8 == 0 {
			fmt.Printf(" and it's divisble by 8\n")
		} else {
			fmt.Printf("\n")
		}
	} else {
		fmt.Printf("%d is an odd number\n", num)
	}
}
