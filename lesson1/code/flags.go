package main

import (
	"flag"
	"fmt"
)

// START OMIT
func main() {
	// A string flag that returns a pointer to the value
	stringFlag := flag.String("name_of_flag",
		"default value",
		"description of how this flag should be used")

	// You can also use the typeVar method where you pass a var to the function
	var intFlag int
	flag.IntVar(&intFlag, "int_flag_name", 0, "description of how this flag should be used")

	// Needs to be called to get the values
	flag.Parse()
	fmt.Println(*stringFlag, intFlag)
}

// END OMIT
