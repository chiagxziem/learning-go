// How to solve the communication between goroutines problem that wait groups can't solve?

package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func sendMessage(ch chan<- string, num int) {
	fmt.Printf("Sending message %d\n", num)

	time.Sleep(time.Second * time.Duration(num)) // simulate some work
	ch <- fmt.Sprintf("✅ Message %d sent!", num)
}

func receiveMessage(msgs <-chan string) {
	fmt.Println("Waiting for message")

	for msg := range msgs {
		fmt.Println("Received:", msg)
	}
}

func main() {
	msgs := make(chan string)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		sendMessage(msgs, 1)
	}()
	go func() {
		defer wg.Done()
		sendMessage(msgs, 2)
	}()

	go func() {
		receiveMessage(msgs)
	}()

	wg.Wait()
	close(msgs) // close the channel to avoid deadlock

	log.Println("Work done")
}
