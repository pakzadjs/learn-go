package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneCHan chan bool) {
	fmt.Println("Hello!", phrase)
	doneCHan <- true
}

func slowGreet(phrase string, doneCHan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneCHan <- true
	close(doneCHan)
}

func main() {
	// dones := make([]chan bool, 4)
	done := make(chan bool)

	// dones[0] = make(chan bool)
	go greet("Nice to meet you!", done)
	// dones[1] = make(chan bool)
	go greet("How are you?", done)
	// dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", done)
	// dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", done)

	// for _, done := range dones {
	// 	<-done
	// }

	for range done {
		// fmt.Println(doneChan)
	}
}
