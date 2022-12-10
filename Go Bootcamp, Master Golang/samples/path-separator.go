package main

import (
	"fmt"
	"path"
)

func main() {
	_, file := path.Split("app/css/main.css")

	fmt.Println(file)
}
