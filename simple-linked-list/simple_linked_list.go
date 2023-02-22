package linkedlist

import "fmt"

// Define the List and Element types here.
// List defined by the pointer to the head element
// Note that the head element is not a real element
type List struct {
	next *Element
	val  int
}

type Element struct {
	next *Element
	val  int
}

// New creates a new list from a slice of values.
func New(sl []int) *List {
	if len(sl) == 0 {
		// return an empty list
		return &List{next: nil}
	}
	// create a new head element
	pList := &List{next: nil}
	for _, s := range sl {
		pList.Push(s)
	}
	return pList
}

// String returns a string representation of the list.
// Implementing the interface Stringer
func (lst *List) String() string {
	out := ""
	if lst.next == nil {
		return out
	}

	curr := lst.next
	for curr.next != nil {
		out += fmt.Sprintf("%v, ", curr.val)
		curr = curr.next
	}
	return out
}

// Size returns the number of elements in the list.
func (lst *List) Size() int {
	//empty list case
	if lst.next == nil {
		return 0
	}
	curr := lst.next
	out := 1
	for curr.next != nil {
		curr = curr.next
		out += 1
	}
	return out
}

// Last returns the last element of the list.
func (lst *List) Last() *Element {
	if lst.next == nil {
		return nil
	}
	curr := lst.next
	for curr.next != nil {
		curr = curr.next
	}
	return curr
}

// Push adds an element to the end of the list.
func (lst *List) Push(elem int) {
	new := &Element{next: nil, val: elem}
	last := lst.Last()
	if last == nil {
		lst.next = new
	} else {

		last.next = new
	}
}

// Pop removes the last element from the list and returns it.
func (lst *List) Pop() (int, error) {
	if lst.next == nil {
		return 0, fmt.Errorf("empty list")
	}
	curr := lst.next
	prev := curr
	for curr.next != nil {
		prev = curr
		curr = curr.next
	}
	out := curr.val
	if curr != prev {
		prev.next = nil
	} else {
		lst.next = nil
	}
	return out, nil
}

// Array returns a slice of the list's elements.
func (lst *List) Array() []int {
	size := lst.Size()
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
	arr := lst.Array()
	out := make([]int, size)
	for i := 0; i < size; i++ {
		out[i] = arr[size-i-1]
	}
	return New(out)
}
