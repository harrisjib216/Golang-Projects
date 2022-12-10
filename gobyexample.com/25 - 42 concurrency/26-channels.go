// channels are pipes which connect concurrent goroutines
// you can send values into channels from one goroutine
// and then receive those values into another goroutine

package main

import "fmt"

// custom example
type person struct {
	name string
	age int
	address string
}

func main() {
	john := make(chan person)

	go func() {
		john <- person{"john", 1, "earth"}
		close(john)
	}()


	for data := range john {
		fmt.Println(data)
	}
}

/* web example
func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
}
*/