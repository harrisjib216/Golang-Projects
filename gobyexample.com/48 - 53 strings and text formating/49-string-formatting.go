package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	// print any type
	fmt.Printf("struct1: %v\n", p)

	// include struct's field names
	fmt.Printf("struct2: %+v\n", p)

	// go syntax representation
	// package.type_name{...values}
	fmt.Printf("struct3: %#v\n", p)

	// type
	fmt.Printf("type: %T\n", p)

	// bool
	fmt.Printf("bool: %b\n", true)

	// int
	fmt.Printf("int: %d\n", 123)

	// binary
	fmt.Printf("bin: %b\n", 14)

	// hex
	fmt.Printf("hex: %x\n", 14)

	// float
	fmt.Printf("float: %f\n", 23.958)

	// scientific notation
	fmt.Printf("scientific notation: %e\n", 12340000.0)
	fmt.Printf("scientific notation: %E\n", 12340000.0)

	// string
	fmt.Printf("string: %s\n", "love")
	fmt.Printf("quoted: %q\n", "love")
	fmt.Printf("hexed: %x\n", "love")

	// rune or character
	fmt.Printf("rune: %c\n", 'a')
	fmt.Printf("character: %c\n", 97)

	// pointer
	fmt.Printf("pointer: %p\n", &p)

	// width
	fmt.Printf("width: |%2d|%2d|\n", 5, 4)

	// percision point
	fmt.Printf("percision point: |%4.4f|%4.5f|\n", 1.2, 1.3)

	// justify left
	fmt.Printf("percision point: |%-4.4f|%-4.5f|\n", 1.2, 1.3)

	// with strings
	fmt.Printf("percision point: |%-4f|%-4f|\n", "testing", "ab")

	// make a string
	str := fmt.Sprintf("sprintf: %s", "some people really like cats")
	fmt.Println(str)

	// output somewhere
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}