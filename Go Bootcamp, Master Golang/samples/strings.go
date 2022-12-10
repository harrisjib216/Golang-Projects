package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	/* String Literals
	- cannot contain multiple lines of text
	- string literals are interpreted
	- when escape sequences are included (\n, \t, etc.)
	go will convert them and add them into our string literal
	*/
	fmt.Print("High performance golang\n")

	/* Raw String Literals
	- can contain multiple lines of text
	- not interpreted
	- go does not "inter-prep(are)" anything inside of the string
	*/
	fmt.Print(`
<html>
	<body>
		"Hello world"
	</body>
</html>
`)

	/* Runes
	- runes are created with single quotes
	- runes are the ascii values for the characters
	- runes are made of bytes
	- string literals are made of runes
	- unicode characters can be 1 to 4 bytes each
	*/
	fmt.Println('a')

	// the number of bytes in a string
	fmt.Println(len("Golang is awesome"))

	// the number of characters in a string
	fmt.Println(utf8.RuneCountInString("Golang is awesome"))
}
