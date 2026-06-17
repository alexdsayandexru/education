package easy

import (
	"fmt"
	"testing"
)

type Result struct {
	buffer []int
}

func Test_144(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 9,
				},
			},
		},
	}

	fmt.Println(preorderTraversal(root))
}

func preorderTraversal(root *TreeNode) []int {
	r := Result{
		buffer: []int{},
	}
	//_preorderTraversal(root, &r)
	_postorderTraversal(root, &r)
	return r.buffer
}

func _preorderTraversal(node *TreeNode, r *Result) {
	if node != nil {
		r.buffer = append(r.buffer, node.Val)

		if node.Left != nil {
			_preorderTraversal(node.Left, r)
		}

		if node.Right != nil {
			_preorderTraversal(node.Right, r)
		}
	}
}

func _postorderTraversal(node *TreeNode, r *Result) {
	if node != nil {
		if node.Left != nil {
			_postorderTraversal(node.Left, r)
			//r.buffer = append(r.buffer, node.Val)
		}
		if node.Right != nil {
			_postorderTraversal(node.Right, r)
			//r.buffer = append(r.buffer, node.Val)
		}
		r.buffer = append(r.buffer, node.Val)
	}
}
