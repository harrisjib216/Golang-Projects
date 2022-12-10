package main

import "fmt"

func addNums(a int, b int) int {
	return a + b
}

func addThreeNums(a, b, c int) int {
	return addNums(a, b) + c
}

func main() {
	onePlusEleven := addNums(1, 11)
	fmt.Println("one plus eleven equals", onePlusEleven)

	result := addThreeNums(addThreeNums(addNums(4, 5), 9, 8), 22, 25)

	fmt.Println("okay here is our result:", result)
}

