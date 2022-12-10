package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Print Your Fullname
//
//  1. Get your name and lastname from the command-line
//  2. Print them using Printf
//
// EXAMPLE INPUT
//  Inanc Gumus
//
// EXPECTED OUTPUT
//  Your name is Inanc and your lastname is Gumus.
// ---------------------------------------------------------

func main() {
	// BONUS: Use a variable for the format specifier
	firstName, lastName := os.Args[1], os.Args[2]

	fmt.Printf("Your name is %s and your last name is %s.\n", firstName, lastName)
}
