package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	str := "sha256 this string"
	hash := sha256.New()

	hash.Write([]byte(str))

	fmt.Println(hash)

	bytes := hash.Sum(nil)

	fmt.Println(str)
	fmt.Println(bytes)
	fmt.Print("%x\n", bytes)
}