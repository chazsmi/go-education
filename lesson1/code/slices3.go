package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s)

	// Add one item to the slice
	s = append(s, 1)
	fmt.Println(s)

	// Add more than on item to the slice
	s = append(s, 2, 3, 4)
	fmt.Println(s)
}
