package w8cye_solution

// package linkedlist
import "errors"

type Element struct {
	next  *Element
	value int
}
type List struct {
	head *Element
	size int
}

func New(values []int) *List {
	list := &List{}
	for _, value := range values {
		list.Push(value)
	}
	return list
}
func (l *List) Size() int {
	return l.size
}
func (l *List) Push(element int) {
	l.head = &Element{value: element, next: l.head}
	l.size++
}
func (l *List) Pop() (int, error) {
	if l.head == nil {
		return 0, errors.New("pop from empty list")
	}
	popValue := l.head.value
	l.head = l.head.next
	l.size--
	return popValue, nil
}
func (l *List) Array() []int {
	size := l.Size()
	array := make([]int, l.Size())
	for n, i := l.head, size-1; n != nil; n, i = n.next, i-1 {
		array[i] = n.value
	}
	return array
}
func (l *List) Reverse() *List {
	var previous *Element
	for current := l.head; current != nil; {
		current.next, previous, current = previous, current, current.next
	}
	l.head = previous
	return l
}
