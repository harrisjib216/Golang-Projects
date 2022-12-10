// timers: execute code in the future once
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting program")

	timer1 := time.NewTimer(2 * time.Second)

	// blocking code here
	<-timer1.C
	fmt.Println("Timer 1 finished after 2 seconds")

	// new example
	timer2 := time.NewTimer(2 * time.Second)
	go func() {
		// blocking code here
		<-timer2.C
		fmt.Println("Timer 2 finished")
	}()

	// stop before the 2 second timer
	stop2 := timer2.Stop()

	// print if we've stopped
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// time.Sleep(2 * time.Second)
}