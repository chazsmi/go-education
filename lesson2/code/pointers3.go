package main

import "fmt"

func main() {
	// START EG OMIT
	a := 1
	b := &a
	fmt.Println(*b) // Would print the value not the pointer
	fmt.Println(b)  // Would print a memory address rather than the value
	// END EG OMIT
}
