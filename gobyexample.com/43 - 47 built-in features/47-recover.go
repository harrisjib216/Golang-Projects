package main

import (
	"fmt"
)

func runPanic() {
	panic("A problem ocurred")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered! Error:", r)
		}
	}()

	runPanic()

	// this code will not run, instead the recover will run
	// idea: learn how to restart programs
	fmt.Println("Code after runPanic()")
}