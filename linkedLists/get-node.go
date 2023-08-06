package linkedList

func (ll *LinkedList[T]) GetNode(index int) T {
	idx := 0
	current := ll.Head
	for current != nil {
		if idx == index {
			return current.Value
		}
		idx++
		current = current.Next
	}
	var zeroValue T
	return zeroValue
}

func (ll *LinkedList[T]) GetNodeRecursive(index int) T {
	current := ll.Head
	return recurseNode(current, 0, index)
}

func recurseNode[T any](head *Node[T], idx int, target int) T {
	if head == nil {
		var zeroValue T
		return zeroValue
	}
	if idx == target {
		return head.Value
	}
	return recurseNode(head.Next, idx+1, target)
}

// func (ll *LinkedList[T]) GetNodeRecursive(index int) *Node[T] {
// 	current := ll.Head
// 	return recurseNode(current, index)
// }

// func recurseNode[T any](head *Node[T], index int) *Node[T] {
// 	if head == nil {
// 		return nil
// 	}
// 	if index == 0 {
// 		return head
// 	}
// 	return recurseNode(head.Next, index-1)
// }
