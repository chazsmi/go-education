package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func pointlessFunc(msg string, quit chan bool) <-chan string {
	c := make(chan string)
	go func() { // Lanunch a go rotutine from the function using an anymonous function
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			quit <- true // We send a signal down the quit channel
		}
	}()
	return c
}

func main() {
	quit := make(chan bool)
	c := pointlessFunc("My names Charlie", quit)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-quit:
			fmt.Println("I'm quiting...")
			return // The main function will finish after this is executed
		}
	}
}

// END OMIT
