package easy

import (
	"fmt"
	"math"
	"testing"
)

func Test_111(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Right: &TreeNode{
				Val: 7,
			},
			Left: &TreeNode{
				Val: 15,
			},
		},
	}
	fmt.Println(minDepth(root))
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getMinDepth(root, 0, math.MaxInt) + 1
}

func getMinDepth(node *TreeNode, level int, min int) int {
	if node.Left != nil {
		min = getMinDepth(node.Left, level+1, min)
	}
	if node.Right != nil {
		min = getMinDepth(node.Right, level+1, min)
	}

	if node.Left == nil && node.Right == nil && level < min {
		min = level
	}
	return min
}
