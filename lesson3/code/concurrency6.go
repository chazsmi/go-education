package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

// START OMIT
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() { // We can have only one go routine in this example
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

// END OMIT
func main() {
	// Call the pointless function to the get the channel back
	c := fanIn(pointlessFunc("Hello charlie"), pointlessFunc("Hello gophers"))
	for i := 0; i < 5; i++ {
		fmt.Printf("Pointless function says: %s\n", <-c)
	}
}
