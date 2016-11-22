package main

import (
	"fmt"
)

func main() {
	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(1, 2))

	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment()) //1
	fmt.Println(increment()) //2

	nextEvent := makeEvenGenerator()
	fmt.Println(nextEvent()) //0
	fmt.Println(nextEvent()) //2
	fmt.Println(nextEvent()) //4
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
