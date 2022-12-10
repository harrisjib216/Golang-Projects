package main

import "fmt"

func passByValue(num int) {
	num += 1
}

func passByPtr(num *int) {
	*num += 1
}

func main() {
	i := 1
	fmt.Println("initially i:", i)
	
	passByValue(i)
	fmt.Println("pass by value\ni:", i)
	
	passByPtr(&i)
	fmt.Println("pass by ptr\ni:", i)

	fmt.Println("address of i:", &i)
}
