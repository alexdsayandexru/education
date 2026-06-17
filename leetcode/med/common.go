package med

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Variant() {
	buffer := []byte("1234567890")
	n := len(buffer)
	for i := range n {
		for j := range i + 1 {
			sub := string(buffer[j : n-i+j])
			fmt.Println(sub)
		}
	}
}

func GetList(values []int) *ListNode {
	var head *ListNode
	var prev *ListNode

	for _, v := range values {
		node := &ListNode{
			Val: v,
		}
		if head == nil {
			head = node
			prev = node
		}
		prev.Next = node
		prev = node
	}
	return head
}

func GetTree() *TreeNode {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	return root
}
