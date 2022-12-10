// a go-routine is a lighweight thread of execution
package main

import (
	"fmt"
	"time"
)

func printThreeTimes(text string) {
	for i := 0; i < 3; i++ {
		fmt.Println(text, ":", i)
	}
}

func main() {
	printThreeTimes("direct")

	go printThreeTimes("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}