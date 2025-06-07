package simple_linked_list

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
	next *Element
	// Tracks the number of elements in the list.
	size int
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
	lst := &List{next: nil, size: 0}
	// pushing elements one at a time is more idiomatic Go way
	for _, val := range sl {
		lst.Push(val)
	}
	return lst
}

// String returns a string representation of the list.
func (lst *List) String() string {
	var builder strings.Builder
	curr := lst.next
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
	curr := lst.next
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
	if lst.next == nil {
		lst.next = newElem
		lst.size++ // increment size for the first element
		return
	}
	last := lst.Last()
	last.next = newElem
	lst.size++ // increment size for each pushed element
}

// Pop removes the last element from the list and returns it.
func (lst *List) Pop() (int, error) {
	if lst.next == nil {
		return 0, fmt.Errorf("empty list")
	}
	curr := lst.next
	if curr.next == nil {
		val := curr.val
		lst.next = nil // remove the only element
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
	curr := lst.next
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
