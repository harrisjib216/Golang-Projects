package main

import (
	"fmt"
	"os"
)

func main() {
	// course
	// get input from the terminal
	// learn about slices

	// task
	if len(os.Args) > 1 {
		fmt.Println("Howdy,", os.Args[1])
	} else {
		fmt.Println("Hey!! You didn't tell me your name!")
	}
}
