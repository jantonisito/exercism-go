package bobahop_solution

import "errors"

type Element struct {
	data int
	next *Element
}

type List struct {
	head *Element
	size int
}

func New(numbers []int) *List {
	output := &List{}

	for _, number := range numbers {
		output.Push(number)
	}

	return output
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	l.head = &Element{element, l.head}
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.size < 1 {
		return 0, errors.New("no elements")
	}
	deadHead := l.head
	l.head = deadHead.next
	deadHead.next = nil
	l.size--
	return deadHead.data, nil
}

func (l *List) Array() []int {
	output := make([]int, l.size)
	for i, head := l.size-1, l.head; i > -1; i, head = i-1, head.next {
		output[i] = head.data
	}
	return output
}

func (l *List) Reverse() *List {
	output := &List{}
	for head := l.head; head != nil; head = head.next {
		output.Push(head.data)
	}
	return output
}
