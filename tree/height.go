package tree

func Height(root *Node[int]) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return 1 + max(Height(root.Left), Height(root.Right))
}
