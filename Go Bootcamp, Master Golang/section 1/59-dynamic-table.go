package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Dynamic Table
//
//  Get the size of the table from the command-line
//    Passing 5 should create a 5x5 table
//    Passing 10 for a 10x10 table
//
// RESTRICTION
//  Solve this exercise without looking at the original
//  multiplication table exercise.
//
// HINT
//  There was a max constant in the original program.
//  That determines the size of the table.
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Give me the size of the table
//
//  go run main.go -5
//    Wrong size
//
//  go run main.go ABC
//    Wrong size
//
//  go run main.go 1
//    X    0    1
//    0    0    0
//    1    0    1
//
//  go run main.go 2
//    X    0    1    2
//    0    0    0    0
//    1    0    1    2
//    2    0    2    4
//
//  go run main.go 3
//    X    0    1    2    3
//    0    0    0    0    0
//    1    0    1    2    3
//    2    0    2    4    6
//    3    0    3    6    9
// ---------------------------------------------------------
var size int
var err error

func printCol(col int) {
	fmt.Printf("%d |", col)
}

func printHeader() {
	fmt.Printf("X |")

	printRow(1)

	for i := 0; i < size*7; i++ {
		fmt.Print("-")
	}

	fmt.Println()
}

func printRow(row int) {
	for col := 0; col <= size; col++ {
		fmt.Printf("%5d", row*col)
	}

	fmt.Println()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Expected: size of the table.")
		return
	}

	size, err = strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Printf("%q is not a valid size.\n", os.Args[1])
		return
	}

	printHeader()

	for i := 0; i <= size; i++ {
		printCol(i)
		printRow(i)
	}

}
