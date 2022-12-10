package main

import (
	"embed"
)

var filestring string

var filebytes []byte

var folder embed.FS

func main() {
	print(filestring)
	print(string(filebytes))

	c1, _ := folder.ReadFile("folder/file1.hash")
	print(string(c1))

	c2, _ := folder.ReadFile("folder/file2.hash")
	print(string(c2))
}