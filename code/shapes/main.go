package main

import "fmt"

type shape interface {
	getArea() float64
}
type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

func main() {
	t := triangle{base: 10, height: 10}
	s := square{sideLength: 10}

	printArea(t)
	printArea(s)
}

func (s square) getArea() float64 { // VERY custom logic => to interface
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 { // VERY custom logic  => to interface
	return 0.5 * t.base * t.height
}

func printArea(s shape) { // VERY common logic => will delegate in interface
	fmt.Println(s.getArea())
}
