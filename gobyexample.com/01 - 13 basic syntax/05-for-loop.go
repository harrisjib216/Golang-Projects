package main

import "fmt"

func main() {
	i := 1
	for i < 4 {
		fmt.Println(i)
		i += 1
	}

	for j := 7; j < 10; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("infinite loop")
		break
	}

	for n := 10; n > -1; n-- {
		if n % 2 == 0 {
			continue
		}

		fmt.Println(n)
	}
}
