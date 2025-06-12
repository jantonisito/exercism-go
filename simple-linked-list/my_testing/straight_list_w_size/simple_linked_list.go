// Package straight_list_w_size implements a simple linked list with size tracking.
// It provides methods to create a list, push and pop elements, retrieve the size of the list
// reverse list in place, return string and array representation of the list.
// Elements are stored sequentioally which poses a problem for the Push and Pop methods.
// -----------------------------------------------------------------------------------------------------------
// Benchmark numbers
// -----------------------------------------------------------------------------------------------------------
// BenchmarkNewList-8               2322774               564.1 ns/op           160 B/op         10 allocs/op
// BenchmarkListSize-8             1000000000               0.5046 ns/op          0 B/op          0 allocs/op
// BenchmarkListPush-8                 1257            920151 ns/op           16000 B/op       1000 allocs/op
// BenchmarkListPop-8                  1593            798561 ns/op               0 B/op          0 allocs/op
// BenchmarkListMixedPopPush-8          404           3017539 ns/op           16000 B/op       1000 allocs/op
// BenchmarkListToArray-8          20443556                56.12 ns/op           80 B/op          1 allocs/op
// BenchmarkListReverse-8           1631755               695.5 ns/op           336 B/op         13 allocs/op
// BenchmarkLongListReverse-8            15          72519120 ns/op          323856 B/op      10003 allocs/op
// -----------------------------------------------------------------------------------------------------------
package straight_list_w_size

// package linked_list // renamed for exercism

import (
	"fmt"
	"strconv"
	"strings"
)

// Define the List and Element types here.

// List defined by the pointer to the first element head.
// This approach allows for easier development of the list.
// due to separation of concern.
type List struct {
	// Pointer to the first element in the list.
	head *Element
	// Tracks the number of elements in the list.
	size int
	// last *Element // pointer to the last element in the list, not used in this implementation
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
	return fmt.Sprintf("%d", e.val)
}

// val returns the value of the element.
func (e *Element) Val() int {
	return e.val
}

// New creates a new linked list from the input slice.
// Note: The input slice `sl` should not be modified externally after calling this function,
// as the function assumes ownership of the slice's data for memory allocation.
// func New(sl []int) *List {
// 	// allocate memory for all elements
// 	elems := make([]Element, len(sl))
// 	// create a new head element
// 	var lst List
// 	for i, val := range sl {
// 		elems[i].val = val
// 		elems[i].next = lst.next
// 		lst.next = &elems[i]
// 	}
// 	return &lst
// }
// Note: The above commented code is a previous version of the New function.
// It creates a new linked list from the input slice by allocating memory for all elements
// and linking them together. The function assumes ownership of the slice's data for memory allocation.
// It is efficient memory allocation and avoids unnecessary copying of data but requires
// New creates a new linked list from the input slice.
// The problem is that the memory for the list is allocated on the stack,
// but Push and Pop methods require the memory to be allocated on the heap.

// function New creates a new linked list from the input slice.
// Note: The input slice `sl` should not be modified externally after calling this function,
// as the function assumes ownership of the slice's data for memory allocation.
func New(sl []int) *List {
	// create a new head element
	lst := &List{head: nil, size: 0}
	// pushing elements one at a time is more idiomatic Go way
	for _, val := range sl {
		lst.Push(val)
	}
	return lst
}

// String returns a string representation of the list.
func (lst *List) String() string {
	var builder strings.Builder
	curr := lst.head
	for curr != nil {
		builder.WriteString(strconv.Itoa(curr.val) + ", ")
		curr = curr.next
	}
	out := builder.String()
	if len(out) > 2 {
		out = out[:len(out)-2] // remove the last ", "
	}
	return out
}

// Size returns the number of elements in the list.
func (lst *List) Size() int {
	return lst.size
}

// Last returns the last element of the list.
func (lst *List) Last() *Element {
	curr := lst.head
	if curr == nil {
		return nil
	}
	for curr.next != nil {
		curr = curr.next
	}
	return curr
}

// Push adds an element to the end of the list.
func (lst *List) Push(elem int) {
	newElem := &Element{val: elem}
	if lst.head == nil {
		lst.head = newElem
		lst.size++ // increment size for the first element
		return
	}
	last := lst.Last()
	last.next = newElem
	lst.size++ // increment size for each pushed element
}

// Pop removes the last element from the list and returns it.
func (lst *List) Pop() (int, error) {
	if lst.head == nil {
		return 0, fmt.Errorf("empty list")
	}
	curr := lst.head
	if curr.next == nil {
		val := curr.val
		lst.head = nil // remove the only element
		lst.size--     // decrement size for the last element
		return val, nil
	}
	// traverse to the second last element
	// to remove the last element
	for curr.next.next != nil {
		curr = curr.next
	}
	toRemove := curr.next
	curr.next = nil
	lst.size-- // decrement size for each popped element

	return toRemove.val, nil
}

// Array returns a slice of the list's elements.
func (lst *List) Array() []int {
	size := lst.size
	if size == 0 {
		return []int{}
	}
	out := make([]int, size)
	curr := lst.head
	for i := 0; i < size; i++ {
		out[i] = curr.val
		curr = curr.next
	}
	return out
}

// Reverse reverses the list.
func (lst *List) Reverse() *List {
	size := lst.Size()
	if size == 0 {
		return lst
	}
	out := make([]int, size)
	arr := lst.Array()
	for i := 0; i < size; i++ {
		out[size-i-1] = arr[i]
	}
	return New(out)
}
