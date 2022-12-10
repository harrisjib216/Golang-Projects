package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 12, height: 13}
	fmt.Println("area by value:", r.area())
	fmt.Println("perimeter by value:", r.perim())

	rPtr := &r
	fmt.Println("area by pointer:", rPtr.area())
        fmt.Println("perimeter by pointer:", rPtr.perim())
}
