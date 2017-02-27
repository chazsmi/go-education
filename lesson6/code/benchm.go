package main

import "strings"

// START OMIT
func AddTo(values []string) string {
	str := "This " + "is " + "my holdall"
	newstring := strings.Join(values, " ")
	return str + newstring
}

// END OMIT
