package linkedList

func (ll *LinkedList[T]) Reversed() *LinkedList[T] {
	var current *Node[T] = ll.Head
	var prev *Node[T] = nil

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	ll.Head = prev

	return ll
}

func RecursiveReverse[T any](head *Node[T], prev *Node[T]) *Node[T] {
	if head == nil {
		return prev
	}
	next := head.Next
	head.Next = prev
	return RecursiveReverse(next, head)
}
