package main

import (
	"fmt"
	"time"
)

var greetings = []string{"Hello", "Hola", "Ni Hao", "Ka", "Sanu"}

func main() {
	//create the channel, we made the channel buffered with the capacity of one by inserting 1 before the closing bracket
	ch := make(chan string, 1)

	// start the greeter to provide a greeting
	go greet(ch)
	// sleep for a while
	time.Sleep(1 * time.Second)
	fmt.Println("main is ready")
	for greeting := range ch {
		// sleep for a while
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Greeting received", greeting)
	}
}

func greet(ch chan string) {
	fmt.Println("Greeter ready")
	for _, g := range greetings {
		ch <- g
	}
	close(ch)
	fmt.Println("Greeting completed")
}
