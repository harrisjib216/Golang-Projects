package main

import "fmt"

func vals(num int) (int, string) {
	var msg string

	switch {
	case num < 0:
		msg = "less than zero"
	case num < 10:
		msg = "has one digit"
	case num < 100:
		msg = "has two digits"
	case num < 1000:
		msg = "has three digits"
	default:
		msg = "has a number of digits"
	}

	return num, msg
}

func main() {
	x, description := vals(24)
	fmt.Printf("%d %s\n", x, description)
}
