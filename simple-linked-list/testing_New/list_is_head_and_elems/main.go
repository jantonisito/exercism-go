package main

import "fmt"

type List struct {
	next *Element
}

type Element struct {
	next *Element
	val  int
}

func New(sl []int) *List {
	if len(sl) == 0 {
		// return an empty list
		return &List{next: nil}
	}
	// allocate memory for all elements
	elems := make([]Element, len(sl))
	pElems := make([]*Element, len(sl))
	for i := range elems {
		pElems[i] = &elems[i]
	}

	// create a new head element
	lst := &List{next: nil}
	for i, s := range sl {
		e := pElems[i]
		(*e).next = lst.next
		(*e).val = s
		lst.next = e
	}
	fmt.Println(lst)
	return lst
}

func (lst *List) Size() int {
	out := 0
	for n := lst.next; n != nil; n = n.next {
		out += 1
	}
	return out
}

func (lst *List) String() string {
	out := ""
	curr := lst.next
	for curr != nil {
		out += fmt.Sprintf("%v, ", curr.val)
		curr = curr.next
	}
	return out
}

func (lst *List) Array() []int {
	size := lst.Size()
	if size == 0 {
		return []int{}
	}
	out := make([]int, size)
	curr := lst.next
	for i := 0; i < size; i++ {
		out[size-i-1] = curr.val
		curr = curr.next
	}
	return out
}

func main() {
	lst := New([]int{1, 2, 3})
	fmt.Println(lst)
	fmt.Println(lst.Array())
}
