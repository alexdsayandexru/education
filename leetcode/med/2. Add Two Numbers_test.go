package med

import (
	"fmt"
	"testing"
)

func Test_2(t *testing.T) {
	l1 := GetList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := GetList([]int{9, 9, 9, 9})
	l3 := addTwoNumbers(l1, l2)
	fmt.Println(l3)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	l3 := _addTwoNumbers(l1, l2)
	return l3
}

func _addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode

	n1 := l1
	n2 := l2
	n3 := l3

	hiN := 0

	for {
		if n1 == nil && n2 == nil {
			break
		}
		d1 := 0
		if n1 != nil {
			d1 = n1.Val
			n1 = n1.Next
		}

		d2 := 0
		if n2 != nil {
			d2 = n2.Val
			n2 = n2.Next
		}

		n31 := &ListNode{
			Val: (d1 + d2 + hiN) % 10,
		}
		if l3 == nil {
			l3 = n31
		}
		if n3 != nil {
			n3.Next = n31
		}
		n3 = n31

		if (d1 + d2 + hiN) >= 10 {
			hiN = 1
		} else {
			hiN = 0
		}
	}

	if hiN == 1 {
		n3.Next = &ListNode{
			Val: hiN,
		}
	}
	return l3
}
