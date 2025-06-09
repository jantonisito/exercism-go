package straight_list_w_size_last_ptr

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
	size  int
	first *Element // pointer to the last element in the reversed list - actual first element
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

// function New creates a new linked list from the input slice.
// Note: The input slice `sl` should not be modified externally after calling this function,
// as the function assumes ownership of the slice's data for memory allocation.
// func New(sl []int) *List {
// 	// create a new head element
// 	lst := &List{first: nil, size: 0, last: nil}
// 	// pushing elements one at a time is more idiomatic Go way
// 	for _, val := range sl {
// 		lst.Push(val)
// 	}
// 	// TODO consider using a more efficient way to update the last pointer
// 	// update the last pointer to point to the last element
// 	if lst.head != nil {
// 		curr := lst.head
// 		for curr.next != nil {
// 			curr = curr.next
// 		}
// 		lst.last = curr
// 	} else {
// 		lst.last = nil // if the list is empty, last should be nil
// 	}
// 	// return the created list
// 	return lst
// }

func New(data []int) *List {
	// create a new head element
	lst := &List{head: nil, size: 0, first: nil}
	size := len(data)
	if size == 0 {
		return lst // return an empty list if the input slice is empty
	}
	if size == 1 {
		// if the input slice has only one element, create a single element list
		lst.head = &Element{val: data[0]}
		lst.first = lst.head // head is the same as first in this case
		lst.size = 1         // set size to 1
		return lst           // return the single element list
	}

	// pushing elements one at a time is more idiomatic Go way
	// for i := (size - 1); i >= 0; i-- {
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
		lst.first = newElem // update pointer to the first element
		lst.size++          // increment size for the first element
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
		lst.head = nil  // remove the only element
		lst.size--      // decrement size for the last element
		lst.first = nil // last should be nil if the list is empty
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
	var next *Element
	for curr != nil {
		next = curr.next // store the next element
		curr.next = prev // reverse the link
		prev = curr      // move prev to the current element
		curr = next
	}
	lst.head, lst.first = lst.first, lst.head // swap first and last pointers
	lst.first.next = nil                      // set the next of the new last element to nil
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
