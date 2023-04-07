package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Math Tables
//
//  Create division, addition and subtraction tables
//
//  1. Get the math operation and
//     the size of the table from the user
//
//  2. Print the table accordingly
//
//  3. Correctly handle the division by zero error
//
//
// BONUS #1
//
//  Use strings.IndexAny function to detect
//    the valid operations.
//
//  Search on Google for: golang pkg strings IndexAny
//
// BONUS #2
//
//  Add remainder operation as well (remainder table using %).
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Usage: [op=*/+-] [size]
//
//  go run main.go "*"
//    Size is missing
//    Usage: [op=*/+-] [size]
//
//  go run main.go "%" 4
//    Invalid operator.
//    Valid ops one of: */+-
//
//  go run main.go "*" 4
//    *    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    1    2    3    4
//    2    0    2    4    6    8
//    3    0    3    6    9   12
//    4    0    4    8   12   16
//
//  go run main.go "/" 4
//    /    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    1    0    0    0
//    2    0    2    1    0    0
//    3    0    3    1    1    0
//    4    0    4    2    1    1
//
//  go run main.go "+" 4
//    +    0    1    2    3    4
//    0    0    1    2    3    4
//    1    1    2    3    4    5
//    2    2    3    4    5    6
//    3    3    4    5    6    7
//    4    4    5    6    7    8
//
//  go run main.go "-" 4
//    -    0    1    2    3    4
//    0    0   -1   -2   -3   -4
//    1    1    0   -1   -2   -3
//    2    2    1    0   -1   -2
//    3    3    2    1    0   -1
//    4    4    3    2    1    0
//
// BONUS:
//
//  go run main.go "%" 4
//    %    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    0    1    1    1
//    2    0    0    0    2    2
//    3    0    0    1    0    3
//    4    0    0    0    1    0
//
// NOTES:
//
//   When running the program in Windows Git Bash, you might need
//   to escape the characters like this.
//
//   This happens because those characters have special meaning.
//
//   Division:
//     go run main.go // 4
//
//   Multiplication and others:
//   (this is also necessary for non-Windows bashes):
//
//     go run main.go "*" 4
// ---------------------------------------------------------

var size int
var err error
var operation string

// table functions
func printCol(col int) {
	fmt.Printf("%d |", col)
}

func printHeader() {
	fmt.Printf("%s |", operation)

	for col := 0; col <= size; col++ {
		fmt.Printf("%5d", col)
	}

	fmt.Println()
}

func printRow(row int) {
	for col := 0; col <= size; col++ {
		fmt.Printf("%5d", operatorMap[operation](row, col))
	}

	fmt.Println()
}

// arithmetic functions
var operatorMap = map[string]func(a, b int) int{
	"+": func(a, b int) int {
		return a + b
	},
	"-": func(a, b int) int {
		return a - b
	},
	"*": func(a, b int) int {
		return a * b
	},
	"/": func(a, b int) int {
		if a == 0 || b == 0 {
			return 0
		}

		return a / b
	},
	"%": func(a, b int) int {
		return a % b
	},
	"^": func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	},
}

func main() {
	// error handling
	if len(os.Args) != 3 {
		fmt.Println("Expected: arithmetic operation and table size.")
		return
	}

	operation = os.Args[1]
	if strings.IndexAny("+-*/%^", operation) == -1 {
		fmt.Printf("%q is not a valid arithmetic operation.\n", os.Args[1])
		return
	}

	size, err = strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("%q is not a valid size.\n", os.Args[2])
		return
	}

	// build dynamic table
	printHeader()
	for i := 0; i <= size; i++ {
		printCol(i)
		printRow(i)
	}
}
