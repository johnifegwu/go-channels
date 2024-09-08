package main

import (
	"fmt"
	"time"
)

var hellos = []string{"Hello", "Hola", "Ni Hao", "Ka", "Sanu"}
var goodbyes = []string{"Goodbye", "Ka chi fo", "Sajien", "Ya ke sua", "Adios"}

func main() {
	//create the channel, we made the channel buffered with the capacity of one by inserting 1 before the closing bracket
	ch := make(chan string, 1)
	ch2 := make(chan string, 1)

	// start the hello goodbye to provide a greeting
	go greet(hellos, ch)
	go greet(goodbyes, ch2)
	// sleep for a while
	time.Sleep(1 * time.Second)
	fmt.Println("main is ready")
	for {
		select {
		case gr, ok := <-ch:
			if !ok {
				ch = nil
				break
			}
			printgreeting(gr)
		case gr2, ok := <-ch2:
			if !ok {
				ch2 = nil
				break
			}
			printgreeting(gr2)
		default:
			return
		}
	}
}

func printgreeting(gr string) {
	// sleep for a while
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Greeting received", gr)
}

func greet(greetings []string, ch chan string) {
	fmt.Println("Greeting ready")
	for _, g := range greetings {
		ch <- g
	}
	close(ch)
	fmt.Println("Greeting completed")
}
