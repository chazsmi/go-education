package main

import (
	"fmt"
	"log"
)

func main() {
	// START ERROR OMIT
	value, err := mypackage.MyFunction()
	if err != nil {
		log.Println(err.Error())
	}
}

// END ERROR OMIT
// START ERROR TYPE OMIT
type MyError struct {
	Msg  string
	File string
	Line int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s:%d: %s", e.File, e.Line, e.Msg)
}

func myFunction() MyError {
	// Error happened
	// ...
	return &MyError{
		Msg:  "Database connection failed",
		File: "database.go",
		Line: 22,
	}
}

// END ERROR TYPE OMIT
