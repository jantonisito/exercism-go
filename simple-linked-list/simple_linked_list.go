// Package reverse_list_w_size implements a simple linked list with a pointer to the last element
// called head and a size field. Elements can be pushed to the  end of the list and popped from the end.
// So the list is stored in the reversed order with last element serving as head.
// This allows for fast push and VERY fast pop operation.
// The list can be reversed efficiently in place so the other methods that need the list in the original order
// (e.g. String or Array) can use the Reverse method.
// -----------------------------------------------------------------------------------------------------------
// Benchmark numbers
// -----------------------------------------------------------------------------------------------------------
// BenchmarkNewList-8               2483992               490.9 ns/op           176 B/op         11 allocs/op
// BenchmarkListSize-8             1000000000               0.4022 ns/op          0 B/op          0 allocs/op
// BenchmarkListPush-8                26324             46629 ns/op           16000 B/op       1000 allocs/op
// BenchmarkListPop-8                302571              3996 ns/op               0 B/op          0 allocs/op
// BenchmarkListMixedPopPush-8        25777             47903 ns/op           16000 B/op       1000 allocs/op
// BenchmarkListToArray-8          14483220                75.71 ns/op           80 B/op          1 allocs/op
// BenchmarkListReverse-8          78223797                13.84 ns/op            0 B/op          0 allocs/op
// BenchmarkLongListReverse-8         79933             13890 ns/op               0 B/op          0 allocs/op
// -----------------------------------------------------------------------------------------------------------
// package reverse_list_w_size

package linkedlist // renamed for exercism

import (
	"fmt"
	"strconv"
	"strings"
)

// Define the List and Element types here.

// List defined by the pointer to the last  element.
// This approach allows for easier development of Push and Pop methods
// due to separation of concern.
type List struct {
	// Pointer to the first element in the list.
	head *Element
	// Tracks the number of elements in the list.
	size int
	// first *Element // pointer to the last element in the reversed list - actual first element
}
type Element struct {
	// Pointer to the next element in the list.
	next *Element
	// Holds the value of the element.
	val int
}

// String implements the Stringer interface for Element.
// It returns a string representation of the element's value.
func (e *Element) String() string {
	return fmt.Sprintf("%d", (*e).val)
}

// val returns the value of the element.
func (e *Element) Val() int {
	return e.val
}

func New(data []int) *List {
	// create a new head element
	lst := &List{head: nil, size: 0}
	size := len(data)
	if size == 0 {
		return lst // return an empty list if the input slice is empty
	}
	if size == 1 {
		// if the input slice has only one element, create a single element list
		lst.head = &Element{val: data[0]}
		lst.size = 1 // set size to 1
		return lst   // return the single element list
	}

	// pushing elements one at a time is more idiomatic Go way
	for i := 0; i < size; i++ {
		lst.Push(data[i])
	}
	return lst
}

// String returns a string representation of the list.
func (lst *List) String() string {
	var builder strings.Builder
	lst.Reverse() // reverse the list to print in original order
	curr := lst.head
	for curr != nil {
		builder.WriteString(strconv.Itoa(curr.val) + ", ")
		curr = curr.next
	}
	out := builder.String()
	if len(out) > 2 {
		out = out[:len(out)-2] // remove the last ", "
	}
	lst.Reverse() // reverse back to original order
	return out
}

// Size returns the number of elements in the list.
func (lst *List) Size() int {
	return lst.size
}

// Last returns the last element of the list.
func (lst *List) Last() *Element {
	return lst.head
}

// Push adds an element to the end of the list.
func (lst *List) Push(elem int) *Element {
	newElem := &Element{val: elem}
	if lst.head == nil {
		lst.head = newElem
		lst.size++ // increment size for the first element
		return newElem
	}
	prevHead := lst.head     // store the previous head
	lst.head = newElem       // update head pointer to the new element
	lst.head.next = prevHead // update the next pointer of the new head to the previous head
	lst.size++               // increment size for each pushed element
	return newElem
}

// Pop removes the last element from the list and returns it.
func (lst *List) Pop() (int, error) {
	if lst.head == nil {
		return 0, fmt.Errorf("empty list")
	}
	var val int
	if lst.size == 1 {
		val = lst.head.val
		lst.head = nil // remove the only element
		lst.size--     // decrement size for the last element
		return val, nil
	}
	toRemove := lst.head     // store the current head to remove it
	val = toRemove.val       // store the value of the first element
	lst.head = toRemove.next // move head to the next element
	toRemove = nil           // clear the removed element to avoid memory leak
	lst.size--               // decrement size for each popped element
	return val, nil
}

// Array returns a slice of the list's elements.
func (lst *List) Array() []int {
	size := lst.size
	if size == 0 {
		return []int{}
	}
	lst.Reverse() // reverse the list to print in original order
	out := make([]int, size)
	curr := lst.head
	for i := 0; i < size; i++ {
		out[i] = curr.val
		curr = curr.next
	}
	lst.Reverse() // reverse the list to print in original order
	return out
}

// Reverse reverses the list.
func (lst *List) Reverse() *List {
	size := lst.size
	if size == 0 || size == 1 {
		// if the list is empty or has only one element, return it as is
		// this is an optimization to avoid unnecessary memory allocation
		return lst
	}
	prev := lst.head
	curr := prev.next
	oldHead := lst.head // store the old head to update it later
	newHead := curr
	var next *Element
	for curr != nil {
		next = curr.next // store the next element
		curr.next = prev // reverse the link
		prev = curr      // move prev to the current element
		newHead = curr   // update newHead to the current element
		curr = next
	}
	lst.head = newHead // set head to the new first element
	oldHead.next = nil // set the next of the old head to nil
	return lst
}
