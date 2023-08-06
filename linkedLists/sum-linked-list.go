package linkedList

func Sum(ll *LinkedList[int]) int {
	current := ll.Head
	res := 0
	for current != nil {
		res += current.Value
		current = current.Next
	}
	return res
}

func SumRecursive(ll *LinkedList[int]) int {
	return recurseSum(ll.Head, 0)
}

func recurseSum(node *Node[int], value int) int {
	if node == nil {
		return value
	}
	value += node.Value
	return recurseSum(node.Next, value)
}
