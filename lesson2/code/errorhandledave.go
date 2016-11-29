package main

// START OMIT

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	e := myVeryWobblyFunction()
	switch err := errors.Cause(e).(type) {
	case MyErrorType:
		// Custom error type print with stack
		myerror := errors.WithStack(e)
		fmt.Println(myerror)
	default:
		// Default error handling
		fmt.Println(err)
		fmt.Printf("%+v", err)
	}
}

func myVeryWobblyFunction() error {
	err := errors.New("this is the cause of the error")
	return errors.Wrap(err, "request failed")
}

// END OMIT
