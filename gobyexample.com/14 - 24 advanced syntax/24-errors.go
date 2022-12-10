package main

import (
	"errors"
	"fmt"
)

type numErr struct {
	num int
	msg string
}

func (ne *numErr) Error() string {
	return fmt.Sprintf("%d - %d", ne.num, ne.msg)
}

func f1(num int) (int, error) {
	if num == 43 {
		return -1, errors.New("Cant use number 43")
	}

	return num + 2, nil
}

func f2(num int) (int, error) {
	if num == 43 {
		return -1, &numErr{num, "can't use this one!"}
	}

	return num + 2, nil
}

func main() {
	for _, n := range []int{7, 43} {
		if res, err := f1(n); err == nil {
			fmt.Println("f1 worked:", res)
		} else {
			fmt.Println("f1 failed:", err)
		}
	}

	for _, n := range []int{7, 43} {
                if res, err := f2(n); err == nil {
                        fmt.Println("f2 worked:", res)
                } else {
                        fmt.Println("f2 failed:", err)
                }
        }

	_, err := f2(43)
	if ne, ok := err.(*numErr); ok {
		fmt.Println(ne.num)
		fmt.Println(ne.msg)
	}
}
