package main

import "fmt"

func main() {
	data := map[string]int{}

	data["k1"] = 7
	data["k2"] = 13

	fmt.Println("map:", data)

	
	val, _ := data["k1"]
	fmt.Println("value:", val)

	
	fmt.Println("length:", len(data))

	
	delete(data, "k2")
	fmt.Println("map after deletion:", data)


	val, exists := data["k2"]
	fmt.Printf("does %v exist?: %v\n", "k2", exists) 


	data2 := map[rune]int{'a':1, 'b':2}
	fmt.Println("delcaration:", data2)
}
