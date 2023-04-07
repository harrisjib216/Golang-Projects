package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Expected 2 arguments: min and max (both integers)")
		return
	}

	min, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a valid integer.\n", os.Args[1])
		return
	}

	max, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("%q is not a valid integer.\n", os.Args[2])
		return
	}

	var sum int
	for min <= max {
		if min%2 != 0 {
			min++
			continue
		}

		fmt.Printf("%d", min)

		if min != max && min+1 != max {
			fmt.Printf(" + ")
		}

		sum += min
		min++
	}

	fmt.Printf(" = %d\n", sum)
}
