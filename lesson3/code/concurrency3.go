package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT

func main() {
	// Make the chan that accepts string
	c := make(chan string)

	// Call the pointless function in a new goroutine
	go pointlessFunc("Hello gophers", &c)

	// Loop, waiting for the function to send data back
	// The main function will be blocked by <-c until it recieves a value
	for i := 0; i < 5; i++ {
		// Receive expression is just a value.
		fmt.Printf("Pointless function says: %s\n", <-c)
	}
}

func pointlessFunc(msg string, c *chan string) {
	for i := 0; i < 5; i++ {
		// Note the change to Sprintf
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

// END OMIT
