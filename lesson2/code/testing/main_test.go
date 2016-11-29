package main

// Your tests are done per package they can be incudled in the same package
// or added in seperate package depending on what your are doung

import (
	"fmt"
	"testing"
)

// START BASIC OMIT
// Naming convenation for tests are Test followed by the function/method you are testing
func TestAverage(t *testing.T) {
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
func TestTableAverage(t *testing.T) {
	// Create a slice of a stuct with ints
	cases := []struct{ A, B, Expected float64 }{
		{1, 2, 1.5},
		{4, 23, 13.5},
		{455, 6667, 3561},
	}
	// Run each case
	for _, tc := range cases {
		actual := Average([]float64{tc.A, tc.B})
		if actual != tc.Expected {
			t.Errorf(
				"got %d, expected %d",
				tc.A, tc.B, actual, tc.Expected,
			)
		}
	}
}

// END TT OMIT

// START TEST1.7 OMIT
func TestTableOneSevenAverage(t *testing.T) {
	// Create a slice of a stuct with ints
	cases := []struct{ A, B, Expected float64 }{
		{1, 2, 1.5},
		{4, 23, 13.5},
		{455, 6667, 356},
	}
	// Run each case
	for _, tc := range cases {
		// The run method allows you run each test in turn and get a fail output
		// for a particluar test
		t.Run(fmt.Sprintf("%f and %f", tc.A, tc.B), func(t *testing.T) { // HL
			actual := Average([]float64{tc.A, tc.B})
			if actual != tc.Expected {
				t.Errorf(
					"got %f, expected %f",
					tc.A, tc.B, actual, tc.Expected,
				)
			}
		}) // HL
	}
}

// END TEST1.7 OMIT

// START BENCH OMIT
// Benchmark functions start with Benchmark not Test.
func BenchmarkAverage(b *testing.B) {
	// Benchmark functions are run several times by the testing package.
	// The value of b.N will increase each time until the benchmark runner
	// is satisfied with the stability of the benchmark
	for n := 0; n < b.N; n++ {
		Average([]float64{1, 2})
	}
}

// END BENCH OMIT

// START TEARDOWN OMIT
func TestFoo(t *testing.T) {
	// <setup code>
	t.Run("A=1", func(t *testing.T) {
		// write the test here...
	})
	t.Run("A=2", func(t *testing.T) {
		// write the test here...
	})
	t.Run("B=1", func(t *testing.T) {
		if !test(foo{B: 1}) {
			t.Fail()
		}
	})
	// <tear-down code>
}

// END TEARDOWN OMIT
