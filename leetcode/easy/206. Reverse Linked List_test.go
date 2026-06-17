package easy

import (
	"fmt"
	"testing"
)

func Test_206(t *testing.T) {
	head := GetList([]int{1, 2, 3, 4, 5})
	tail := reverseList(head)
	fmt.Println(tail)
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var tail *ListNode
	_reverseList(nil, head, &tail)
	return tail
}

func _reverseList(prev *ListNode, node *ListNode, tail **ListNode) {
	if node.Next != nil {
		_reverseList(node, node.Next, tail)
	}
	if *tail == nil {
		*tail = node
	}
	node.Next = prev
}
