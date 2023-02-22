package linkedlist

import "fmt"

// Define the List and Element types here.
// List defined by the pointer to the head element
// Note that the head element is not a real element
type List struct {
	next *List
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
func (l *List) String() string {
	out := ""
	curr := l
	if curr == nil {
		return out
	}
	for curr.next != nil {
		out += fmt.Sprintf("%v, ", curr.val)
		curr = curr.next
	}
	return out
}

// Size returns the number of elements in the list.
func (l *List) Size() int {
	curr := l
	if l == nil {
		return 0
	}
	out := 0
	for curr.next != nil {
		curr = curr.next
		out += 1
	}
	return out
}

// Last returns the last element of the list.
func (l *List) Last() *List {
	curr := l
	if curr == nil {
		return curr
	}
	for curr.next != nil {
		curr = curr.next
	}
	return curr
}

// Push adds an element to the end of the list.
func (l *List) Push(element int) {
	last := l.Last()
	if last == nil {
		l.val = element
		return
	}
	new := &List{next: nil, val: element}
	last.next = new
}

// Pop removes the last element from the list and returns it.
func (l *List) Pop() (int, error) {
	if l.next == nil {
		return 0, fmt.Errorf("empty list")
	}
	curr := l
	prev := curr
	for curr.next != nil {
		prev = curr
		curr = curr.next
	}
	out := curr.val
	prev.next = nil
	return out, nil
}

// Array returns a slice of the list's elements.
func (l *List) Array() []int {
	size := l.Size()
	if size == 0 {
		return []int{}
	}
	out := make([]int, size)
	curr := l.next
	for i := 0; i < size; i++ {
		out[i] = curr.val
		curr = curr.next
	}
	return out
}

// Reverse reverses the list.
func (l *List) Reverse() *List {
	size := l.Size()
	if size == 0 {
		return l
	}
	arr := l.Array()
	out := make([]int, size)
	for i := 0; i < size; i++ {
		out[i] = arr[size-i-1]
	}
	return New(out)
}
