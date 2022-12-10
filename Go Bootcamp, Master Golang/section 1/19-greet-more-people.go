package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Greet More People
//
// RESTRICTIONS
//  1. Be sure to match the expected output below
//  2. Store the length of the arguments in a variable
//  3. Store all the arguments in variables as well
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {
	// TYPE YOUR CODE HERE
	c := len(os.Args)

	if c == 1 {
		fmt.Println("Ya gotta add some people to the arguments list.")
		return
	}

	if c == 2 {
		fmt.Println("There is one person!")
		defer fmt.Println("Nice to you meet you.")
	} else {
		fmt.Println("There are", c-1, "people!")
		defer fmt.Println("Nice to meet you all.")
	}

	for i := 1; i <= c-1; i++ {
		fmt.Println("Hello great", os.Args[i], "!")
	}

	// BONUS #1:
	// Observe the error if you pass less then 3 arguments.
	// Search on the web how to solve that.
}
