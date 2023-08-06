package tree

import "math"

func min(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	currMin := numbers[0]
	for _, n := range numbers {
		if n < currMin {
			currMin = n
		}
	}
	return currMin
}

func MinValue(root *Node[int]) int {
	if root == nil {
		return math.MaxInt32
	}
	leftMin := MinValue(root.Left)
	rightMin := MinValue(root.Right)
	return min(leftMin, rightMin, root.Value)
}
