package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	msg := os.Args[1]
	length := utf8.RuneCountInString(msg)

	// append as many exclamation marks as the length of the string
	marks := strings.Repeat("!", length)

	// capitalize the paramter
	msg = strings.ToUpper(marks + msg + marks)

	fmt.Println(msg)
}
