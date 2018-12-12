package main

import "fmt"

type circle struct {
	radius float64
}

type rectangle struct {
	length  float64
	breadth float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	area := 3.14157 * c.radius * c.radius
	return area
}

func (r rectangle) area() float64 {
	area := r.length * r.breadth
	return area
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	fmt.Println("HI")
	a1 := rectangle{}
	a1.length, a1.breadth = 10, 10
	b1 := circle{}
	b1.radius = 10
	fmt.Println(b1.area(), a1.area())
}
