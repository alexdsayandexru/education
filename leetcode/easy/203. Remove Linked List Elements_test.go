package easy

import (
	"fmt"
	"testing"
)

func Test_203(t *testing.T) {
	values := []int{1, 2, 2, 1}
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
	head2 := removeElements(head, 2)
	fmt.Println(head2)
}

func removeElements(head *ListNode, val int) *ListNode {
	var prev *ListNode
	top := head
	node := head

	for node != nil {
		if node.Val == val {
			if node == top {
				top = node.Next
				prev = node
				node = node.Next
			} else {
				prev.Next = node.Next
				node = node.Next
			}
		} else {
			prev = node
			node = node.Next
		}
	}
	return top
}
