package straight_list_w_size_last_ptr

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
	size int
	last *Element // pointer to the last element in the list
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

	// pushing elements one at a time is more idiomatic Go way
	for i := 0; i < (size - 1); i++ {
		lst.Push(data[i])
	}
	lst.last = lst.Push(data[size-1]) // push the last element separately
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
	newElem := &Element{val: elem}
	if lst.head == nil {
		lst.head = newElem
		lst.last = newElem // update last pointer to the first element
		lst.size++         // increment size for the first element
		return newElem
	}
	lst.last.next = newElem // update last pointer to the new element
	lst.last = newElem      // update the last pointer to the new element
	lst.size++              // increment size for each pushed element
	return newElem
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
		lst.last = nil // last should be nil if the list is empty
		return val, nil
	}
	// traverse to the second last element
	// to remove the last element
	for curr.next.next != nil {
		curr = curr.next
	}
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
