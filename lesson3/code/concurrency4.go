package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func pointlessFunc(msg string) <-chan string {
	c := make(chan string)
	go func() { // Lanunch a go rotutine from the function using an anymonous function
		for i := 0; i < 5; i++ {
			// Note the change to Sprintf
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	// Call the pointless function to the get the channel back
	c := pointlessFunc("Hello gophers")
	for i := 0; i < 5; i++ {
		fmt.Printf("Pointless function says: %s\n", <-c)
	}
}

// END OMIT
