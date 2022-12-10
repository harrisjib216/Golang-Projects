package main

import "fmt"

func main() {
	nums := []int{2,3,4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for index, num := range nums {
		if num == 3 {
			fmt.Println("3 is found at index:", index)
		}
	}

	keyValStore := map[string]string{"a": "apple","b": "banana"}

	// range with key and value
	for key, val := range keyValStore {
		fmt.Printf("%s -> %s\n", key, val)
	}

	// range just with key
	for key := range keyValStore {
		fmt.Println("key:", key)
	}

	// byte index, rune
	for byteIndex, runeVal := range "golang" {
		fmt.Println(byteIndex, runeVal)
	}
}
