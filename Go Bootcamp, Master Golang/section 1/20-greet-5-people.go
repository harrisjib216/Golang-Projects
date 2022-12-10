package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Greet 5 People
//
//  Greet 5 people this time.
//
//  Please do not copy paste from the previous exercise!
//
// RESTRICTION
//  This time do not use variables.
//
// INPUT
//  bilbo balbo bungo gandalf legolas
//
// EXPECTED OUTPUT
//  There are 5 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Hello great gandalf !
//  Hello great legolas !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {
	// TYPE YOUR CODE HERE
	if len(os.Args) == 1 {
		fmt.Println("You gotta add some people to the list")
	} else if len(os.Args) == 2 {
		fmt.Println("There is one person here!")
		fmt.Println("Hi great", os.Args[1], "!")
		fmt.Println("Nice to meet you.")
	} else {
		fmt.Println("There are 5 people!")
		fmt.Println("Hi great", os.Args[1], "!")
		fmt.Println("Hi great", os.Args[2], "!")
		fmt.Println("Hi great", os.Args[3], "!")
		fmt.Println("Hi great", os.Args[4], "!")
		fmt.Println("Hi great", os.Args[5], "!")
		fmt.Println("Nice to meet you all.")
	}
}
