package main

import "fmt"

// START OMIT
// This function now takes a pointer
func answer(yes *bool) {
	*yes = true
}

func main() {
	yes := false
	// We pass a pointer to the function rather than the value itself
	answer(&yes)
	fmt.Println(yes) // yes is now true
}

// END OMIT
