package tree

import "github.com/DavidEsdrs/ads/queue"

func TreeInclude(t *Tree[string], target string) bool {
	queue := queue.Queue[*Node[string]]{}
	queue.Push(t.Root)
	for queue.Size > 0 {
		current := *queue.Shift()
		if current.Value == target {
			return true
		}
		if current.Right != nil {
			queue.Push(current.Right)
		}
		if current.Left != nil {
			queue.Push(current.Left)
		}
	}
	return false
}

func TreeRecursive(t *Node[string], target string) bool {
	if t == nil {
		return false
	}
	if t.Value == target {
		return true
	}
	return TreeRecursive(t.Left, target) || TreeRecursive(t.Right, target)
}
