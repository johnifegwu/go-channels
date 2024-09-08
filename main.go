package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	//create the channel, we made the channel buffered with the capacity of one by inserting 1 before the closing bracket
	ch := make(chan string, 1)

	// start the greeter to provide a greeting
	go greet(ch)
	// sleep for a while
	time.Sleep(5 * time.Second)
	fmt.Println("main is ready")
	greeting, ok := <-ch
	if !ok {
		fmt.Println("Channel is closed")
		log.Fatal("Channel is closed")
	}
	// sleep for a while
	time.Sleep(2 * time.Second)
	fmt.Println("Greeting received")
	fmt.Println(greeting)
}

func greet(ch chan string) {
	fmt.Printf("Greeter ready.\nGreeter waiting to send greeting...\n")
	ch <- "Hello, world"
	fmt.Println("Greeting completed")
}
