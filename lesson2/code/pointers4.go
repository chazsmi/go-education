package main

import "fmt"

func main() {
	// START NEW OMIT
	d := new(int)
	fmt.Println(d)  // Would print a memory address rather than the value
	fmt.Println(*d) // Would print the value
	// END NEW OMIT
}
