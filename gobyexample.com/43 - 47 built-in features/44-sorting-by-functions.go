// how to use a custom function for sorting
package main

import (
	"fmt"
	"sort"
)

type byLength []string

// todo: try pointer
func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "mango", "lime", "watermelon", "banana"}
	sort.Sort(byLength(fruits))
	fmt.Println("Sorted fruits by word length:", fruits)
}