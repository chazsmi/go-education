package main

func main() {
	var inbasket bool
	oranges := 1
	if inbasket {
		// Do stuff
	}
	if oranges != 1 {
		// Do stuff
	}

	// Short hand
	if i := 1; i < oranges {
		// Do stuff
		// We can only use i here though
	}
	// Cant use i here
}
