// defer is used to ensure a code is executed
// last in its scope's execution
package main

import (
	"fmt"
	// "log"
	"os"
)

func main() {
	file := createFile("created-file.txt")
	defer closeFile(file)

	writeFile(file)
}

func createFile(f string) *os.File {
	fmt.Println("creating", f)

	file, err := os.Create(f)

	if err != nil {
		panic(err)
	}

	return file
}

func closeFile(f *os.File) {
	fmt.Println("closing")

	err := f.Close()

	if err != nil {
		panic(err)
	}
}

func writeFile(f *os.File) {
	fmt.Println("writing")

	fmt.Fprintln(f, "testing")
}