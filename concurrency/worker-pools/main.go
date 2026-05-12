package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		time.Sleep(time.Second) //simulate work
		fmt.Println("Worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	// This creates a fixed number of workers

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// start the workers
	for i := range 3 {
		go worker(i+1, jobs, results)
	}

	// send in the jobs
	for i := range numJobs {
		jobs <- i + 1
	}
	close(jobs)

	// get out the results
	for range numJobs {
		<-results
	}
	close(results)

	fmt.Println("all jobs have been processed!")
}
