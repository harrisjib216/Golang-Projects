package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "food", "a string")

	numPtr := flag.Int("num", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string variable")

	flag.Parse()

	fmt.Println("word", *wordPtr)
	fmt.Println("num", *numPtr)
	fmt.Println("fork", *forkPtr)
	fmt.Println("svar", svar)
	fmt.Println("tail", flag.Args())
}