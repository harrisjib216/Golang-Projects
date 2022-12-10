// wait for multiple goroutines to finish
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // call done, but for each goroutine we spawn, signals the end of a goroutine
					// and therefore decreases our wg counter
	
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)

    fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	// wg.Add(5) // directly set number of goroutines for which to wait

	for i := 1; i <= 5; i++ {
		wg.Add(1) // increment to number how many goroutines for which we will wait

		i := i // create new variable to prevent closure error

		go func() {
			worker(i, &wg)
		}() // can also use parameter syntax
	}

	wg.Wait() // wait to terminate until our counter reaches 0
}