package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int

	if a := os.Args; len(a) != 2 {
		fmt.Println("Give me a number")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		fmt.Printf("Error: Cannot convert %q\n", a[1])
	} else {
		fmt.Printf("Successfully converted %d\n", n)
	}

	fmt.Println("N now equals", n)
}
