// closing a channel means no more values
// will be sent on it
package main

import "fmt"

func main() {
	jobs := make(chan int, 5) // remove the 5 for a race condition
	done := make(chan bool)

	go func() {
		for {
			j, more := <- jobs

			if more {
				fmt.Println("Received job", j)
			} else {
				fmt.Println("Received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j < 4; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)

	fmt.Println("sent all jobs")

	// wait for the done value
	<-done // remove this line to exit early and not wait for our jobs to finish
}