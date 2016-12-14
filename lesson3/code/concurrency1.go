package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT

func main() {
	// We have added the go statement before calling the function
	// This will lanch the pointlessFunc in a new goroutine
	go pointlessFunc("Hello gophers")

}

func pointlessFunc(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

// END OMIT
