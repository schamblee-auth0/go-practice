package main

import "fmt"

type shape interface {
	printArea() float64
}

type triangle struct {
	sideLength int
}
type square struct {
	sideLength int
}

func main() {
	triangle := triangle{
		sideLength: 2,
	}
	square := square{
		sideLength: 2,
	}

	getArea(triangle)
	getArea(square)
}

func getArea(s shape) {
	fmt.Println(s.printArea())
}

func (t triangle) printArea() float64 {
	return float64((t.sideLength * t.sideLength) / 2)
}

func (s square) printArea() float64 {
	return float64(s.sideLength * s.sideLength)
}
