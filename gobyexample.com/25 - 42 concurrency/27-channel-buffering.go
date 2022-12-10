// by default channels are unbuffered which means
// channels will only accept sends: channel <-
// if there is a corresponding receive: <- channel
// buffered channels accept a limited number of values
// without a corresponding receiver for those values

package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	// send to first
	messages <- "buffered"

	// send to second
	messages <- "channel"

	// receive first
	fmt.Println(<-messages)

	// receive second
	fmt.Println(<-messages)	
}