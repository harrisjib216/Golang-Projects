package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Vowel or Consonant
//
//  Detect whether a letter is vowel or consonant.
//
// NOTE
//  y or w is called a semi-vowel.
//  Check out: https://en.oxforddictionaries.com/explore/is-the-letter-y-a-vowel-or-a-consonant/
//
// HINT
//  + You can find the length of an argument using the len function.
//
//  + len(os.Args[1]) will give you the length of the 1st argument.
//
// BONUS
//  Use strings.IndexAny function to detect the vowels.
//  Search on Google for: golang pkg strings IndexAny
//
// Furthermore, you can also use strings.ContainsAny. Check out: https://golang.org/pkg/strings/#ContainsAny
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a letter
//
//  go run main.go hey
//    Give me a letter
//
//  go run main.go a
//    "a" is a vowel.
//
//  go run main.go y
//    "y" is sometimes a vowel, sometimes not.
//
//  go run main.go w
//    "w" is sometimes a vowel, sometimes not.
//
//  go run main.go x
//    "x" is a consonant.
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]

	if len(args) == 0 || utf8.RuneCountInString(args[0]) > 1 {
		fmt.Println("Give me a letter")
		return
	}

	letter := args[0]

	if strings.IndexAny(letter, "aeiou") == 0 {
		fmt.Printf("%s is a vowel.\n", letter)
	} else if letter == "y" {
		fmt.Println("y is sometimes a vowel")
	} else {
		fmt.Printf("%s is a consonant.\n", letter)
	}
}
