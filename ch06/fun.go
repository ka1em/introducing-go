package main

import (
	"fmt"
)

func main() {
	xs := []float64{1, 2, 3}

	total := 0.0
	for _, v := range xs {
		total += v
	}
	fmt.Println(len(xs))
	fmt.Println(total / float64(len(xs)))
}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}
