package main

import "testing"

// START OMIT
func BenchmarkAddTo(b *testing.B) {
	items := []string{
		"trainers",
		"jacket",
		"tshirt",
		"socks",
		"pants",
		"goggles",
		"swimcap",
	}
	b.ReportAllocs()
	// You can reset the timer to exclude any setup time in the benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AddTo(items)
	}
}

// END OMIT
