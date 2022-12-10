// when using channels as function parameters
// we can specify if a channel is meant to only
// send or only receive values

package main

import "fmt"

// one channel to send: ping
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// one channel to send: ping
// one channel to receive: pong
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 2)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	/* how it works
	1. make a ping channel to send
	2. make a pong channel to receive
	3. ping function sends "msg" to "pings" channel
	4. pong function receives "msg"/data from "pings" channel and then
	   sends that message into the "pongs" channel
	5. lastly we consume/receive the data/"msg" from the pongs channel
}