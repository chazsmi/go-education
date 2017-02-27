package main

// START OMIT
func AddTo(values []string) string {
	strs := []byte("This ")
	strs = append(strs, []byte("is ")...)
	strs = append(strs, []byte(" my holdall")...)
	for _, b := range values {
		strs = append(strs, []byte(b)...)
	}
	return string(strs)
}

// END OMIT
