package main

// Your tests are done per package they can be incudled in the same package
// or added in seperate package depending on what your are doung

import "testing"

// START BASIC OMIT
// Naming convenation for tests are Test followed by the function/method you are testing
func TestAverage(t *testing.T) r
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		// Testing package has a range of functions allowing you
		// to fail or pass tests with messages
		t.Error("Expected 1.5, got ", v)
	}
}

// END BASIC OMIT
// START TT OMIT
func TestAdd(t testing.T) {
	// Create a slice of a stuct with ints
	cases := []struct{ A, B, Expected int }{
		{1, 2, 2},
		{3, 1, 4},
		{1, -3, 2},
	}
	// Run each case
	for k, tc := range cases {
		actual := tc.A + tc.B
		if actual != expected {
			t.Errorf(
				"%d + %d = %d, expected %d",
				tc.A, tc.B, actual, tc.Expected,
			)
		}
	}
}

// END TT OMIT
