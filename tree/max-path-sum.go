package tree

import "math"

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func MaxPathSum(root *Node[int]) int {
	if root == nil {
		return math.MinInt32
	}
	if root.Left == nil && root.Right == nil {
		return root.Value
	}
	leftMax := MaxPathSum(root.Left)
	rightMax := MaxPathSum(root.Right)
	return root.Value + max(leftMax, rightMax)
}
