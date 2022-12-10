package main

// ---------------------------------------------------------
// EXERCISE: Simplify the Leap Year
//
//  1. Look at the solution of "the previous exercise".
//
//  2. And simplify the code (especially the if statements!).
//
// EXPECTED OUTPUT
//  It's the same as the previous exercise.
// ---------------------------------------------------------

/*
original
import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a year number")
		return
	}

	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a valid year.\n", os.Args[1])
		return
	}

	var leap bool
	if year%400 == 0 {
		leap = true
	} else if year%100 == 0 {
		leap = false
	} else if year%4 == 0 {
		leap = true
	}

	if leap {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}
*/

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a year number")
	} else if year, err := strconv.Atoi(os.Args[1]); err != nil {
		fmt.Printf("%q is not a valid year.\n", os.Args[1])
	} else if year%4 == 0 {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}
