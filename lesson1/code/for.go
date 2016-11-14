package main

import "fmt"

func main() {
	// Bog standard for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	// Go's version of a while loop
	sumWhile := 1
	for sumWhile < 1000 {
		sumWhile += sumWhile
	}

	// Go's foreach loop
	// Don't worry I will explain what a map is later, but you can probably guess from the code
	var array = map[string]int{"apples": 1, "oranges": 20, "bananas": 4}
	for key, value := range array {
		fmt.Println(key)
		fmt.Println(value)
	}
	// end
}
