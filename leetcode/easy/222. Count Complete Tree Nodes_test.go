package easy

import (
	"fmt"
	"testing"
)

func Test_222(t *testing.T) {
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
	fmt.Println(countNodes(root))
}

func countNodes(root *TreeNode) int {
	var count int
	_countNodes(root, &count)
	return count
}

func _countNodes(node *TreeNode, count *int) {
	if node != nil {
		if node.Left != nil {
			_countNodes(node.Left, count)
		}
		if node.Right != nil {
			_countNodes(node.Right, count)
		}
		*count = *count + 1
	}
}
