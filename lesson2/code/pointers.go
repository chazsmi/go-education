package main

import "fmt"

// START OMIT
func answer(yes bool) {
	yes = true
}

func main() {
	yes := false
	answer(yes)
	fmt.Println(yes) // yes is still false
}

// END OMIT
