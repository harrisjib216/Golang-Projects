package main

import (
	"fmt"
	"time"
)

// id, receiver, sender
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for w := 1; w <= numJobs; w++ {
		go worker(w, jobs, results)
	}

	for a := 1; a <= numJobs; a++ {
        <-results
    }
}