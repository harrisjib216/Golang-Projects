package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("rand: 0 <= n < 100", rand.Intn(100))
	fmt.Println(rand.Intn(100))

	fmt.Println("random float:", rand.Float64())

	fmt.Println((rand.Float64()*5)+5)
	fmt.Println(rand.Float64() * 5 + 5, "\n")

	src1 := rand.NewSource(time.Now().UnixNano())
	rnd1 := rand.New(src1)

	fmt.Println(rnd1.Intn(100))
	fmt.Println(rnd1.Intn(100), "\n")

	// default num generator is deterministic
	// produces same sequence of numbers with
	// the same seed
	src2 := rand.NewSource(42)
	rnd2 := rand.New(src2)
	fmt.Println(rnd2.Intn(100))
	fmt.Println(rnd2.Intn(100))
	fmt.Println(rnd2.Intn(100), "\n")

	src3 := rand.NewSource(42)
	rnd3 := rand.New(src3)
	fmt.Println(rnd3.Intn(100))
	fmt.Println(rnd3.Intn(100))
	fmt.Println(rnd3.Intn(100))
}