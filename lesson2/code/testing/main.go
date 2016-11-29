package main

import "fmt"

func main() {
	fmt.Println(Average([]float64{455, 6667}))
}

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}
