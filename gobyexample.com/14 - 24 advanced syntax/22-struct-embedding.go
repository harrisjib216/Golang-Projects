package main

import "fmt"

type base struct {
	rep int
}

func (b base) describe() string {
	return fmt.Sprintf("base %v", b.rep)
}


type container struct {
	base
	str string
}


func main() {
	c := container{
		base: base{
			rep: 2,
		},
		str: "base 2 system",
	}

	fmt.Printf("c = {rep: %v, str: %v}\n", c.rep, c.str)
	fmt.Println("also rep:", c.base.rep)
	fmt.Println("c description:", c.describe())


	type describer interface {
		describe() string
	}

	var d describer = c
	fmt.Println("describer:", d.describe())
}
