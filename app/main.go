package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	ts := triangle{height: 5.0, base: 10.0}
	ss := square{sideLength: 10}

	printArea(ts)
	printArea(ss)
}

func (ts triangle) getArea() float64 {
	return 0.5 * ts.height * ts.base
}

func (ss square) getArea() float64 {
	return ss.sideLength * ss.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
