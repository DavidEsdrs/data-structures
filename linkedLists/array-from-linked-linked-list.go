package linkedList

func (ll *LinkedList[T]) ToArray() []*Node[T] {
	arr := []*Node[T]{}
	current := ll.Head
	for current != nil {
		arr = append(arr, current)
		current = current.Next
	}
	return arr
}

func (ll *LinkedList[T]) GetValues() []T {
	arr := ll.ToArray()
	res := []T{}
	for _, n := range arr {
		res = append(res, n.Value)
	}
	return res
}

func (ll *LinkedList[T]) GetValuesRecursively() []T {
	values := []T{}
	recurse(ll.Head, &values)
	return values
}

func recurse[T any](head *Node[T], values *[]T) {
	if head == nil {
		return
	}
	*values = append(*values, head.Value)
	recurse(head.Next, values)
}
