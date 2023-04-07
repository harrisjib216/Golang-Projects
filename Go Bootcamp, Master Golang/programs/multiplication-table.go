package main

import "fmt"

const size int = 5

func printRow(multiplier int) {
	for j := 0; j <= size; j++ {
		fmt.Printf("%5d", multiplier*j)
	}

	fmt.Println("")
}

func printCol(col int) {
	fmt.Printf("%d |", col)
}

func main() {
	fmt.Print("X |")

	printRow(1)

	for i := 0; i <= 32; i++ {
		fmt.Print("-")
	}

	println()

	for i := 0; i <= size; i++ {
		printCol(i)
		printRow(i)
	}
}
