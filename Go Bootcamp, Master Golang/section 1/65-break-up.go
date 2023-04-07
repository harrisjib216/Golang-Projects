package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
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
	for {

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

		if min >= max {
			break
		}
	}

	fmt.Printf(" = %d\n", sum)
}
