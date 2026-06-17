package easy

import (
	"fmt"
	"testing"
)

func Test_226(t *testing.T) {
	root := GetTree()
	root = invertTree(root)
	fmt.Println(root)
}

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		if root.Left != nil {
			invertTree(root.Left)
		}
		if root.Right != nil {
			invertTree(root.Right)
		}
		root.Left, root.Right = root.Right, root.Left
	}
	return root
}
