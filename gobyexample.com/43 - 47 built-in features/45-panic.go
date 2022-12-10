// we use panic to fail fast on errors which shouldn't have
// ocurred
package main

import "os"

func main() {
	panic("a problem ocurred")

	_, err := os.Create("./tmp/file")
	if err != nil {
		panic(err)
	}
}