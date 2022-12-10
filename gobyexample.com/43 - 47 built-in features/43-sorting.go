// go's sort package provides sorting for built-in types and
// user defined types
package main

import (
	"fmt"
	"sort"
)

func main() {
	// strings
	strs := []string{"c", "a", "d", "g", "b", "e"}
	fmt.Println("Before sort:", strs)
	sort.Strings(strs)
	fmt.Println("After sort:", strs, "\n")

	// ints
	ints := []int{5,11,94,33,22,1}
	fmt.Println("Before sort:", ints)
	sort.Ints(ints)
	fmt.Println("After sort:", ints, "\n")

	// check if slice is already sorted
	intsAreSorted := sort.IntsAreSorted(ints)
	fmt.Println("is the ints slice sorted?:", func() string {
		if intsAreSorted {
			return "yes!"
		}

		return "no!"
	}())
}