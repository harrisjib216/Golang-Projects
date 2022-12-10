package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("initialize an empty array", a)

	a[4] = 100
	fmt.Println("set a value in empty array", a)
	fmt.Println("get a value from an array", a[4])
	
	fmt.Println("get the length of an array", len(a))

	b := [5]int{1,2,3,4,5}
	fmt.Println("declare an aray", b)

	var twoD [2][3]int
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[0]); j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
}
