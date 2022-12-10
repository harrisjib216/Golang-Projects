package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextIntFn := intSeq()

	fmt.Println(nextIntFn())
	fmt.Println(nextIntFn())
	fmt.Println(nextIntFn())
	
	nextIntFn = intSeq()
	fmt.Println(nextIntFn())
}
