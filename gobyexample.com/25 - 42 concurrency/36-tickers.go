// tickers allow us to do something repeatedly at regular intervals
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second / 2)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Ticking at", t)
			}
		}
	}()

	// duration of code
	time.Sleep(1600 * time.Millisecond)

	// clean up
	ticker.Stop()

	// stop the ticker by sending true
	done <- true

	fmt.Println("Ticker finished")
}