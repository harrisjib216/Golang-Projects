package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print Your Name and LastName
//
//  Print your name and lastname using Printf
//
// EXPECTED OUTPUT
//  My name is Inanc and my lastname is Gumus.
//
// BONUS
//  Store the formatting specifier (first argument of Printf)
//    in a variable.
//  Then pass it to printf
// ---------------------------------------------------------

func main() {
	// BONUS: Use a variable for the format specifier

	// fmt.Printf("?", ?, ?)
	firstName := "James"
	lastName := "Bond"

	fmt.Printf("My name? It's %s. %[2]s %[1]s", lastName, firstName)
}
