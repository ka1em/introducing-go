package main

import (
	"fmt"
	"introducing-go/ch09/math"
)

func Add(a, b int) int {
	return a + b
}

func main() {

	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
	fmt.Println(Add(1, 2))
}
