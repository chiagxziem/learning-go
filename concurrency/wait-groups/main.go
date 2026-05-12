package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID int
}

type Result struct {
	JobID int
	Value string
}

func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {
		fmt.Println("Worker", id, "processing job", j.ID)
		time.Sleep(1 * time.Second) //simulate work
	}
}

func main() {
	jobs := make(chan Job)
	wg := sync.WaitGroup{}

	workerCount := 4
	jobCount := 20

	for i := range workerCount {
		wg.Add(1)
		go worker(i+1, jobs, &wg)
	}

	for i := range jobCount {
		jobs <- Job{
			ID: i + 1,
		}
	}
	close(jobs)

	wg.Wait()
	fmt.Println("all jobs have been processed!")
}
