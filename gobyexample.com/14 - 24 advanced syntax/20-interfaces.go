package main

import (
	"fmt"
	"math"
)

// geomtry interface
type geometry interface {
	area() float64
	perim() float64
}

func measure(g geometry) {
	fmt.Println("item:", g)
        fmt.Println("area:", g.area())
        fmt.Println("perim:", g.perim())
}

// rectangle
type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// circle
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 { // circumference
	return 2 * math.Pi * c.radius
}


// program
func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	
	measure(r)
	measure(c)
}
