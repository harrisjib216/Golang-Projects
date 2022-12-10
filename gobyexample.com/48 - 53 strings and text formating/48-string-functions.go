package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func main() {
	p("test	-	Contains	- ester", strings.Contains("test", "ester"))
	p("test	-	Count		- tr", strings.Count("test", "tr"))
	p("test	-	HasPrefix	- love", strings.HasPrefix("test", "love"))
	p("test	-	HasSuffix	- mouse", strings.HasSuffix("test", "mouse"))
	p("test	-	Index		- t", strings.Index("test", "t"))
	p("test	-	Index		- te", strings.Index("test", "te"))
	p("test	-	Index		- es", strings.Index("test", "es"))

	p("")
	p("Join:		", strings.Join([]string{"g", "o", "l", "a", "n", "g"}, ""))
	p("golang:		", strings.Split("golang", ""))

	
	p("a	-	Repeat		- 20", strings.Repeat("a", 20))
	p("pool	-	Replace		- 0", strings.Replace("pool", "o", "0", -1))
	p("pool	-	Replace		- 0", strings.Replace("pool", "o", "0", 1))
	p("TEST	-	ToLower:	", strings.ToLower("TEST"))
	p("test	-	ToUpper:	", strings.ToUpper("test"))
}