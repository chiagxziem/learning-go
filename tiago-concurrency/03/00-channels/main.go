package main

import (
	"fmt"
	"time"
)

func sendMessage(num int, msgChan chan<- string) {
	fmt.Printf("Sending message %d\n", num)

	time.Sleep(time.Second * time.Duration(num)) // simulate some work

	msg := fmt.Sprintf("✅ Message %d sent!", num)

	msgChan <- msg
}

func receiveMessage(msgs <-chan string) {
	fmt.Println("Waiting for messages")

	for msg := range msgs {
		fmt.Println("Received:", msg)
	}
}

func main() {
	msgChan := make(chan string)

	go sendMessage(2, msgChan)
	go sendMessage(3, msgChan)

	go receiveMessage(msgChan)

	close(msgChan)
}
