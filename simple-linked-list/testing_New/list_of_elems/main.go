// implements a linked list of integers - head of the list is element
package main

import (
	"fmt"
)

type Element struct {
	next *Element
	val  int
}

func NewList(array []int) *Element {
	lst := &Element{next: nil, val: 0}
	curr := lst
	elements := make([]Element, len(array))
	for i, val := range array {
		e := elements[i]
		curr.val = val
		curr.next = &e
		curr = &e

	}
	return lst
}
func (lst *Element) String() string {
	out := ""
	curr := lst
	for curr.next != nil {
		out += fmt.Sprintf("%v->", curr.val)
		curr = curr.next
	}
	if len(out) > 0 {
		out = out[:len(out)-2] // Remove the trailing ", "
	}

	return out
}

func main() {
	lst := NewList([]int{1, 2, 3})
	fmt.Println(lst)
}
