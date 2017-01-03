package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// START OMIT
	var mutex = sync.Mutex{}
	data := []int{1, 4, 222, 33}

	for c := 0; c < 10; c++ {
		go func(data []int) {
			for r := 0; r < 10; r++ {
				mutex.Lock()
				rand.Seed(time.Now().Unix())
				n := rand.Intn(len(data)-1) + 1
				data[n]++
				mutex.Unlock()
				fmt.Printf("I added to %d! \n", n)
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			}
		}(data)
	}
	// Delay the finishing off the main function to allow go routines to run
	time.Sleep(time.Second * 2)
	fmt.Println("%+v", data)

	// END OMIT
}
