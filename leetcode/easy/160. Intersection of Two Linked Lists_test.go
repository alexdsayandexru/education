package easy

import (
	"fmt"
	"testing"
)

func Test_160(t *testing.T) {
	c1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
		},
	}

	a1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  1,
				Next: c1,
			},
		},
	}

	b1 := &ListNode{
		Val:  3,
		Next: c1,
	}

	fmt.Println(getIntersectionNode(a1, b1))
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	m := make(map[*ListNode]bool)
	for {
		if headA != nil {
			if m[headA] {
				return headA
			} else {
				m[headA] = true
				headA = headA.Next
			}
		}
		if headB != nil {
			if m[headB] {
				return headB
			} else {
				m[headB] = true
				headB = headB.Next
			}
		}
		if headA == nil && headB == nil {
			break
		}
	}
	return nil
}
