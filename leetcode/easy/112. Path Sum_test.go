package easy

import (
	"fmt"
	"testing"
)

func Test_122(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	fmt.Println(hasPathSum(root, 18))
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	ok := false
	isPathSum(root, 0, targetSum, &ok)
	return ok
}

func isPathSum(node *TreeNode, sum int, targetSum int, ok *bool) {
	if *ok {
		return
	}
	if node.Left != nil {
		isPathSum(node.Left, sum+node.Val, targetSum, ok)
	}
	if node.Right != nil {
		isPathSum(node.Right, sum+node.Val, targetSum, ok)
	}

	if node.Left == nil && node.Right == nil && node.Val+sum == targetSum {
		*ok = true
	}
}
