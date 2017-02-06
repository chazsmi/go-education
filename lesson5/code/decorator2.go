package main

import (
	"log"
	"net/http"
)

// START OMIT
// We define what a decorator to our client will look like
type Decorator func(Client) Client

// Decorate will decorate a client with a slice of passed decorators
func Decorate(c Client, ds ...Decorator) Client {
	decorated := c
	for _, decorate := range ds {
		decorated = decorate(decorated)
	}
	return decorated
}

// END OMIT
// START CLIENT OMIT
type Client interface {
	Do(*http.Request) (*http.Response, error)
}

type ClientFunc func(*http.Request) (*http.Response, error)

func (f ClientFunc) Do(r *http.Request) (*http.Response, error) {
	return f(r)
}

// END CLIENT OMIT
// START LOGGER OMIT
func Log(l *log.Logger) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (res *http.Response, err error) {
			l.Printf("%s: %s %s", r.UserAgent(), r.Method, r.URL)
			return c.Do(r)
		})
	}
}

// END LOGGER OMIT
