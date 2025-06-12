// Package: my_testing/straight_list_w_size_last_and_penultimate_ptr implements a linked list of integers
// with a size, last, and penultimate pointers. Pointers to the last and penultimate elements
// are used to optimize the performance of the list operations, such as push and pop.
// Penultimate pointer is cached to avoid traversing the list to find the second last element.
// Once the cached penultimate elemenet is usec cache is invalidated and recalculated either on
// the next push or pop operation. This approach allows for efficient list operations especially
// in cases of mixed push and pop operations, where the penultimate element is frequently accessed
// and the cache is invalidated but trivially recalculated in the next push operation.
// -----------------------------------------------------------------------------------------------------------
// Benchmark numbers
// -----------------------------------------------------------------------------------------------------------
// BenchmarkNewList-8               2412472               480.6 ns/op           208 B/op         11 allocs/op
// BenchmarkListSize-8             1000000000               0.3250 ns/op          0 B/op          0 allocs/op
// BenchmarkListPush-8                32265             39139 ns/op           16000 B/op       1000 allocs/op
// BenchmarkListPop-8                  2374            435944 ns/op               0 B/op          0 allocs/op
// BenchmarkListMixedPopPush-8        29137             44736 ns/op           16000 B/op       1000 allocs/op
// BenchmarkListToArray-8          23009090                51.52 ns/op           80 B/op          1 allocs/op
// BenchmarkListReverse-8          85336698                13.01 ns/op            0 B/op          0 allocs/op
// BenchmarkLongListReverse-8         84051             12331 ns/op               0 B/op          0 allocs/op
// -----------------------------------------------------------------------------------------------------------
package straight_list_w_size_last_and_penultimate_ptr

// package linked_list // renamed for exercism

import (
	"fmt"
	"strconv"
	"strings"
)

// Define the List and Element types here.

// List defined by the pointer to the first element.
// This approach allows for easier development of the list.
// due to separation of concern.
type List struct {
	// Pointer to the first element in the list.
	head *Element
	// Tracks the number of elements in the list.
	size                int
	last                *Element // pointer to the last element in the list
	penultimate         *Element // pointer to the penultimate element in the list
	isValidaPenultimate bool
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

// function New creates a new linked list from the input slice.
func New(data []int) *List {
	// create a new head element
	lst := &List{head: nil, size: 0, last: nil}
	size := len(data)
	if size == 0 {
		return lst // return an empty list if the input slice is empty
	}
	if size == 1 {
		// if the input slice has only one element, create a single element list
		lst.head = &Element{val: data[0]}
		lst.last = lst.head // last is the same as first in this case
		lst.size = 1        // set size to 1
		return lst          // return the single element list
	}
	if size == 2 {
		// if the input slice has only one element, create a single element list
		lst.head = &Element{val: data[0]}
		lst.last = &Element{val: data[1], next: nil}
		lst.head.next = lst.last       // link the first element to the last
		lst.penultimate = lst.head     // penultimate is the first element
		lst.isValidaPenultimate = true // penultimate is valid
		lst.size = 2                   // set size to 2
		lst.last.next = nil            // last element's next is nil
		return lst                     // return the single element list
	}
	// pushing elements one at a time is more idiomatic Go way
	for i := 0; i < (size - 2); i++ {
		lst.Push(data[i])
	}
	lst.penultimate = lst.Push(data[size-2]) // push the last element separately
	lst.last = lst.Push(data[size-1])        // push the last element
	lst.isValidaPenultimate = true           // penultimate is valid
	lst.size = size                          // set size to the number of elements in the list
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
	return lst.last
}

// Push adds an element to the end of the list.
func (lst *List) Push(elem int) *Element {
	penultimate := lst.last // store the current last element as penulti
	newElem := &Element{val: elem}
	if lst.head == nil {
		lst.head = newElem
		lst.last = newElem // update last pointer to the first element
		lst.penultimate = nil
		lst.isValidaPenultimate = false // penultimate is not valid for the first element
		lst.size++                      // increment size for the first element
		return newElem
	}
	lst.last.next = newElem        // update last pointer to the new element
	lst.last = newElem             // update the last pointer to the new element
	lst.size++                     // increment size for each pushed element
	lst.penultimate = penultimate  // update penultimate to the previous last element
	lst.isValidaPenultimate = true // penultimate is valid
	return newElem
}

// Pop removes the last element from the list and returns it.
func (lst *List) Pop() (int, error) {
	if lst.head == nil {
		return 0, fmt.Errorf("empty list")
	}
	// if the list has only one element, remove it
	curr := lst.head
	var val int
	if lst.size == 1 {
		val = curr.val
		lst.head = nil                  // remove the only element
		lst.size--                      // decrement size for the last element
		lst.last = nil                  // last should be nil if the list is empty
		lst.penultimate = nil           // penultimate should also be nil if the list is empty
		lst.isValidaPenultimate = false // penultimate is not valid if the list is empty
		// return the value of the removed element
		return val, nil
	}
	// if the list has two elements, remove the last element
	if lst.size == 2 {
		val = curr.next.val
		curr.next = nil // remove the last element
		lst.head = curr // remove the only element
		lst.size--      // decrement size for the last element
		lst.last = curr
		lst.penultimate = nil           // penultimate should also be nil if the list has only one element
		lst.isValidaPenultimate = false // penultimate is not valid if the list has only one element
		// return the value of the removed element
		return val, nil
	}
	if lst.isValidaPenultimate {
		val = lst.last.val              // store the value of the last element to return it later
		lst.last = lst.penultimate      // update last to the penultimate element
		lst.last.next = nil             // set the next of the new last element to nil
		lst.size--                      // decrement size for each popped element
		lst.penultimate = nil           // reset penultimate to nil
		lst.isValidaPenultimate = false // penultimate is not valid after popping the last element
		return val, nil
	}
	// traverse to the second last element
	// to remove the last element
	for curr.next.next != nil {
		curr = curr.next
	}
	lst.penultimate = curr         // update penultimate to the second last element
	lst.isValidaPenultimate = true // penultimate is valid after popping the last element

	toRemove := curr.next
	curr.next = nil
	lst.last = curr // update the last pointer to the second last element
	lst.size--      // decrement size for each popped element

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
	size := lst.size
	if size == 0 || size == 1 {
		// if the list is empty or has only one element, return it as is
		// this is an optimization to avoid unnecessary memory allocation
		return lst
	}
	prev := lst.head
	curr := prev.next
	var next *Element
	for curr != nil {
		next = curr.next // store the next element
		curr.next = prev // reverse the link
		prev = curr      // move prev to the current element
		curr = next
	}
	lst.head, lst.last = lst.last, lst.head // swap first and last pointers
	lst.last.next = nil                     // set the next of the new last element to nil
	return lst
}

// // Reverse reverses the list in place. Copilot version
// func (lst *List) Reverse() *List {
// 	if lst.size == 0 || lst.size == 1 {
// 		return lst
// 	}
// 	var prev *Element
// 	curr := lst.head
// 	lst.last = lst.head // after reversal, the first becomes the last
// 	for curr != nil {
// 		next := curr.next
// 		curr.next = prev
// 		prev = curr
// 		curr = next
// 	}
// 	lst.head = prev
// 	return lst
// }
