package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 22
	return &p
}

func main() {
	// without fields
	fmt.Println(person{"Bob", 20})
	
	// with fields
	fmt.Println(person{name: "Alice", age: 30})

	// without all fields
	fmt.Println(person{name: "Fred"})
	
	// getting address of item
	fmt.Println(&person{name: "Ann", age: 40})
	
	// this code breaks because there are too few values to declare
	// this struct item
	// fmt.Println(person{"Testing"})

	fmt.Println(newPerson("Jack"))

	sean := person{name: "Sean", age: 50}
	fmt.Println(sean.name)

	seanPtr := &sean
	fmt.Println(seanPtr.age)

	seanPtr.age = 53
	fmt.Println(seanPtr.age)
	
	fmt.Println(&seanPtr.name)
}
