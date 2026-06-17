package easy

import (
	"fmt"
	"math"
	"testing"
)

func Test_110(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 30,
				Left: &TreeNode{
					Val: 40,
				},
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
	}
	fmt.Println(isBalanced(root))
}

func isBalanced(root *TreeNode) bool {
	ok := true
	areBalanced(root, &ok)
	return ok
}

func areBalanced(root *TreeNode, ok *bool) {
	if root == nil || !*ok {
		return
	}
	areBalanced(root.Left, ok)
	areBalanced(root.Right, ok)

	if *ok {
		*ok = _isBalanced(root)
	}
}

func _isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	dl := -1
	if root.Left != nil {
		dl = getDepth(root.Left, 0, 0)
	}
	dr := -1
	if root.Right != nil {
		dr = getDepth(root.Right, 0, 0)
	}

	return math.Abs(float64(dl-dr)) <= 1
}

func getDepth(node *TreeNode, level int, max int) int {
	if node.Left != nil {
		max = getDepth(node.Left, level+1, max)
	}
	if node.Right != nil {
		max = getDepth(node.Right, level+1, max)
	}
	if level > max {
		max = level
	}
	return max
}
