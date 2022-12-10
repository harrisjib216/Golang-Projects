package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	// tests if a pattern matches a string
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// optimized regexp
	r, _ := regexp.Compile("p([a-z]+)ch")

	// match test again
	fmt.Println(r.MatchString("peach"))

	// finds the match for regexp
	fmt.Println(r.FindString("peach punch"))

	// finds the first match, and returns
	// start and end indicies of the match
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// info about all of the pattern matching
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// return indicies^
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// return all of the matches
	fmt.Println(r.FindAllString("cow peach punch pinch", -1))

	// print all submatch indicies
	fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch", -1))

	// print first 2 or X matches
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// determine if match exists in byte array
	fmt.Println(r.Match([]byte("peach")))

	// print the regular expression
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	// replace all ocurrences
	fmt.Println(r.ReplaceAllString("a peach peach", "<fruit>"))

	//
	input := []byte("a peach")
	out := r.ReplaceAllFunc(input, bytes.ToUpper)
	fmt.Println(string(out))
}