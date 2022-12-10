package main

import "fmt"

func main() {
	mass := 123.2
	acceleration := 22

	force := mass * float64(acceleration)

	fmt.Println(force)
}
