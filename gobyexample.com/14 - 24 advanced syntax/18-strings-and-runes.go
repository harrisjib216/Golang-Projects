package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s string = "สวัสดี"

	fmt.Println("len of raw bytes:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println("\nRune count:", utf8.RuneCountInString(s))

	for index, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, index)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	}

	if r == 'ส' {
		fmt.Println("found so sua")
	}
}
