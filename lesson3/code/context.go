package main

import "time"

// START OMIT
// This is the context interface
type Context interface {
	// This will return the time if set that the context has to complete
	Deadline() (deadline time.Time, ok bool)
	// Returns a channel thats closed when the context is canceled or times out
	Done() <-chan struct{}
	// Indictes why a conext was canceled
	Err() error
	// Returns the value by the key set in context or nil
	Value(key interface{}) interface{}
}

// END OMIT
