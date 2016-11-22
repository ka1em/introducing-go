package main

import (
	"fmt"
)

func main() {
	type Circle struct {
		x, y, z float64
	}

	var c Circle
	/*
		d := new(Circle)

		e := Circle{
			x: 0,
			y: 0,
			z: 5}
		f := Circle{0, 0, 5}
	*/
	g := &Circle{0, 0, 5}
	fmt.Println(g.x, g.y, g.z)

	fmt.Println(c.x, c.y, c.z)
}
