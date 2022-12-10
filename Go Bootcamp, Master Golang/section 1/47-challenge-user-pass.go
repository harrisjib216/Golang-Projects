package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

func main() {
	params := os.Args[1:]

	if len(params) < 2 {
		fmt.Println("Usage: [username] [password]")
	} else if params[0] == "jack" && params[1] == "1888" {
		fmt.Printf("Access granted to %q.\n", params[0])
	} else if params[0] == "jack" && params[1] != "1888" {
		fmt.Printf("Invalid password for %q.\n", params[0])
	} else {
		fmt.Printf("Access denied for %q.\n", params[0])
	}
}
