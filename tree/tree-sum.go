package tree

import "github.com/DavidEsdrs/ads/queue"

func SumNodes(root *Node[int]) int {
	if root == nil {
		return 0
	}
	return SumNodes(root.Left) + SumNodes(root.Right) + root.Value
}

func Sum(root *Node[int]) int {
	if root == nil {
		return 0
	}
	res := 0
	q := queue.Queue[*Node[int]]{}
	q.Unshift(&root)
	for q.Size > 0 {
		current := *q.Shift()
		res += current.Value
		if current.Left != nil {
			q.Unshift(&current.Left)
		}
		if current.Right != nil {
			q.Unshift(&current.Right)
		}
	}
	return res
}
