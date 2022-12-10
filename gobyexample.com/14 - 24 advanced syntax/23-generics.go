package main

import "fmt"

// linked list
type element[T any] struct {
	next *element[T]
	val T
}

type List[T any] struct {
	head, tail *element[T]
	// todo: add length
}

func (l *List[T]) Push(v T) {
	if l.tail == nil {
		l.head = &element[T]{val: v}
		l.tail = l.head
	} else {
		l.tail.next = &element[T]{val: v}
		l.tail = l.tail.next
	}

	// todo: add length
}

func (l *List[T]) GetValues() []T {
	var items []T
	
	for curr := l.head; curr != nil; curr = curr.next {
		items = append(items, curr.val)
	}

	return items
}

func GetMapKeys[K comparable, V any](m map[K]V) []K {
	a := make([]K, 0, len(m))

	for k := range m {
		a = append(a, k)
	}
	
	return a
}

func main() {
	m := map[int]string{1: "2", 2: "4", 4: "8"}

	fmt.Println("keys of m:", GetMapKeys(m))

	test := GetMapKeys[int, string](m)
	fmt.Println("keys of test:", test)

	l := List[int]{}
	l.Push(10)
	l.Push(100)
	l.Push(1000)

	fmt.Println("list:", l.GetValues())
}
