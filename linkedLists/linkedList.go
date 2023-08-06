package linkedList

import "fmt"

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T any] struct {
	Head *Node[T]
	Size uint
}

func NewLinkedList[T any]() *LinkedList[T] {
	ll := LinkedList[T]{Head: nil, Size: 0}
	return &ll
}

func FromArray[T any](arr []T) *LinkedList[T] {
	ll := LinkedList[T]{Size: 0}
	for _, n := range arr {
		ll.Append(n)
	}
	return &ll
}

// Append the value to the tail of the linked list
func (ll *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}
	ll.Size++
	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// Append the value as the head of the linked list
func (ll *LinkedList[T]) Prepend(value T) {
	newNode := &Node[T]{Value: value}
	newNode.Next = ll.Head
	ll.Head = newNode
	ll.Size++
}

// Remove an item at given index - returns an error if it don't succed
func (ll *LinkedList[T]) RemoveAtIndex(index int) error {
	if index > int(ll.Size) {
		return fmt.Errorf("index out of range - given index " + fmt.Sprint(index) + " to list of Size " + fmt.Sprint(ll.Size))
	}
	if index < 0 || ll.Head == nil {
		return fmt.Errorf("invalid index or empty list")
	}
	if index == 0 {
		ll.Head = ll.Head.Next
		ll.Size--
		return nil
	}
	current := ll.Head
	count := 0
	var prev *Node[T] = nil
	for current != nil {
		if count == index {
			prev.Next = current.Next
			ll.Size--
			return nil
		}
		prev = current
		current = current.Next
		count++
	}
	return fmt.Errorf("index out of range - given index " + fmt.Sprint(index) + " to list of Size " + fmt.Sprint(ll.Size))
}

func (ll *LinkedList[T]) TraverseIterate() {
	curr := ll.Head
	for curr != nil {
		fmt.Printf("%v", curr.Value)
		curr = curr.Next
	}
}

func (ll *LinkedList[T]) Traverse() {
	next(ll.Head)
}

func next[T any](n *Node[T]) {
	if n != nil {
		fmt.Printf("%v", n.Value)
		next(n.Next)
	}
}

type Person struct {
	Name string
	Age  uint
}

func Execute() {
	dav := &Person{Name: "David", Age: 18}
	arr2 := []*Person{
		dav,
		{Name: "Ryan", Age: 21},
		{Name: "Ju", Age: 20},
		{Name: "Faker", Age: 20},
	}
	arr := []*Person{
		{Name: "Roberto", Age: 46},
		{Name: "Rayssa", Age: 19},
		{Name: "Leonardo", Age: 26},
		{Name: "Tupac", Age: 28},
		{Name: "Tyler", Age: 26},
	}

	ll := FromArray(arr)
	ll2 := FromArray(arr2)

	res := ZipperList(ll, ll2)

	for _, n := range res.ToArray() {
		println(n.Value.Name)
	}
}
