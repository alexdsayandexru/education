package easy

import (
	"fmt"
	"math"
	"testing"
)

func Test_141(t *testing.T) {
	head3 := &ListNode{
		Val: 3,
	}
	head2 := &ListNode{
		Val: 2,
	}
	head0 := &ListNode{
		Val: 0,
	}
	head4 := &ListNode{
		Val:  -4,
		Next: head2,
	}
	head3.Next = head2
	head2.Next = head0
	head0.Next = head4

	fmt.Println(hasCycle(head3))
}

func hasCycle(head *ListNode) bool {
	var ok bool
	isCycle(head, &ok)
	return ok
}

func isCycle(node *ListNode, ok *bool) {
	if node == nil || node.Next == nil {
		*ok = false
	} else {
		if node.Val != math.MaxInt {
			node.Val = math.MaxInt
			isCycle(node.Next, ok)
		} else {
			*ok = true
		}
	}
}
