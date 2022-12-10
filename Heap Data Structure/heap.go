package main

import "fmt"

type MinHeap struct {
	data []int
}

// Insert
// add an element to the end of the heap
// heapify that element, place it where it needs to be
func (h *MinHeap) Insert(val int) {
	h.data = append(h.data, val)

	child := len(h.data) - 1
	parent := int((child - 1) / 2)

	for parent >= 0 && h.data[parent] > h.data[child] {
		h.Swap(parent, child)
		child = parent
		parent = int((child - 1) / 2)
	}
}

// Extract
// return the head/largest key
// place the last element in the heap as the head
// heapify the tree
func (h *MinHeap) Pop() (int, bool) {
	if len(h.data) == 0 {
		return -1, false
	}

	// extract first
	min := h.data[0]

	// truncate
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	last := len(h.data) - 1

	// heapify down
	parent := 0
	left := parent*2 + 1
	right := parent*2 + 2

	for left <= last {
		fmt.Println(parent, left, right)
		if (right > last || h.data[left] < h.data[right]) && h.data[left] < h.data[parent] {
			// left child is smaller
			h.Swap(parent, left)
			parent = left
		} else if right <= last && h.data[right] < h.data[parent] {
			// right child is smaller
			h.Swap(parent, right)
			parent = right
		} else {
			break
		}

		left = parent*2 + 1
		right = parent*2 + 2
	}

	return min, true
}

func (h *MinHeap) Swap(one, two int) {
	h.data[one], h.data[two] = h.data[two], h.data[one]
}

func (h *MinHeap) Size() int {
	return len(h.data)
}

func main() {
	mh := MinHeap{}

	mh.Insert(10)
	mh.Insert(5)
	mh.Insert(100)
	mh.Insert(3)
	mh.Insert(2)
	fmt.Println(mh)

	mh.Pop()
	fmt.Println(mh)

	kp := []int{}
	for i := 0; i < mh.Size(); i++ {
		kp = append(kp)
	}
}
