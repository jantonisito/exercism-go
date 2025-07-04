// implements a linked list of integers - head of the list is element
package main

import (
	"fmt"
	// sl "simple-linked-list/my_testing/straight_list_w_size_last_ptr"
	// sl "simple-linked-list/my_testing/straight_list_w_size" // Uncomment this line to use the previous version of the linked list
	// sl "simple-linked-list/my_testing/straight_list_w_size_last_and_penultimate_ptr"
	// sl "simple-linked-list/my_testing/reverse_list_w_size_last_ptr"
	sl "simple-linked-list/my_testing/reverse_list_w_size"
)

func main() {
	lst := sl.New([]int{1, 2, 3})
	fmt.Println(lst)
	fmt.Println("Size of the list:", lst.Size())
	fmt.Println("Last element of the list:", lst.Last())
	lst.Push(5)
	fmt.Println("After pushing 5:", lst)
	fmt.Println("Size of the list after pushing 5:", lst.Size())
	val, _ := lst.Pop()
	fmt.Println("Popped value:", val)
	fmt.Println("List after popping last element:", lst)

	arr := lst.Array()
	fmt.Println("List converted to array:", arr)

	list := sl.New([]int{})
	list.Push(1)
	list.Push(2)
	// list.Push(3)
	elem := list.Push(3)
	fmt.Printf("Element pushed: %v\n", elem)
	fmt.Printf("Size of list from [] after 3 calls to Push(): %d\n", list.Size())
	fmt.Printf("List after pushing 1, 2, 3: %s\n", list)
	fmt.Println("Last element of the list after pushing 1, 2, 3:", list.Last())
	rev := list.Reverse()
	fmt.Printf("List after reverse: %s\n", rev)
	fmt.Println("Original list after reverse:", list)

	list.Pop()
	list.Pop()
	list.Pop()
	fmt.Printf("Size of list from [] after 3 calls to Pop(): %d\n", list.Size())
	fmt.Printf("List after 3 calls to Pop: %s\n", list)
	fmt.Println("Last element of the list after  3 pops :", list.Last())
}
