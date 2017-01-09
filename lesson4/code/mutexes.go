package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// START SETUP OMIT
type Cache struct {
	sync.Mutex // Protects the store below
	Number     map[int]string
}

func New() *Cache {
	return &Cache{
		Number: make(map[int]string),
	}
}

func (c *Cache) Get(key int) string {
	c.Lock()
	defer c.Unlock()
	return c.Number[key]
}

func (c *Cache) Set(key int, value string) {
	c.Lock()
	defer c.Unlock()
	c.Number[key] = value
}

// END SETUP OMIT

// START OMIT
func main() {
	ca := New()
	for c := 0; c < 10; c++ {
		go func(cache *Cache) {
			for r := 0; r < 10; r++ {
				s1 := rand.NewSource(time.Now().UnixNano())
				r1 := rand.New(s1)
				key := r1.Intn(2)
				cache.Set(key, "charlie")
				fmt.Printf("I added to %d! \n", cache.Number[key])
				time.Sleep(time.Duration(1 * time.Millisecond))
			}
		}(ca)
	}
	// Delay the finishing off the main function to allow go routines to run
	time.Sleep(time.Second * 2)
	fmt.Println("%+v", ca)
}

// END OMIT
