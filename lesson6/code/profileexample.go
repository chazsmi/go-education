package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.MemProfile).Stop()
	port := flag.String("port", "8000", "Server port")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		selectFoods()
	})

	// Starts the web server
	http.ListenAndServe(":"+*port, nil)
}

func selectFoods() string {
	food := []string{
		"Salad",
		"chips",
		"ketchup",
		"icecream",
		"apples",
		"bananas",
		"curry",
		"milk",
		"eggs",
	}
	feedme := make(chan string)
	for i := 0; i < 100; i++ {
		go func() {
			feedme <- food[random(1, len(food))]
		}()
	}
	c := 0
	yourfood := []string{}
	for {
		yourfood = append(yourfood, <-feedme)
		c++
		if len(yourfood) > 2 {
			return fmt.Sprintf("You have %s", yourfood)
		}
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}
