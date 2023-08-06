package linkedList

func Exists[T comparable](ll *LinkedList[T], target T) bool {
	current := ll.Head
	for current != nil {
		if current.Value == target {
			return true
		}
		current = current.Next
	}
	return false
}

func ExistsRecursive(head *Node[int], target int) bool {
	if head == nil {
		return false
	}
	if head.Value == target {
		return true
	}
	return ExistsRecursive(head.Next, target)
}
