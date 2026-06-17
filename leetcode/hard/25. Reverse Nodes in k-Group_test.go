package hard

import (
	"testing"
)

func Test_25(t *testing.T) {
	head := GetList([]int{1, 2, 3, 4, 5})

	reverseKGroup(head, 3)
	/*f1, l1, ok1 := split(head, 1)
	fmt.Println(f1, l1, ok1)
	f2, l2, ok2 := split(l1, 1)
	fmt.Println(f2, l2, ok2)
	f3, l3, ok3 := split(l2, 1)
	fmt.Println(f3, l3, ok3)*/

	/*var top *ListNode
	reverse(head, head.Next, &top)
	fmt.Println(top)*/
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	ok := true
	var result *ListNode
	var top *ListNode
	var last *ListNode

	curr := head
	for curr != nil && ok {
		head, curr, ok = split(curr, k)
		if ok {
			reverse(head, head.Next, &top)
			if result == nil {
				result = top
			}
			if last != nil {
				last.Next = top
			}
			last = head
		} else {
			last.Next = head
		}
	}
	return result
}

func split(first *ListNode, k int) (*ListNode, *ListNode, bool) {
	curr := first
	for i := 0; i < k-1; i++ {
		if curr.Next != nil {
			curr = curr.Next
		} else {
			return first, curr, false
		}
	}
	last := curr.Next
	curr.Next = nil
	return first, last, true
}

func reverse(curr *ListNode, next *ListNode, top **ListNode) {
	if curr != nil && next != nil {
		curr.Next = nil
		reverse(next, next.Next, top)
		next.Next = curr
	} else {
		*top = curr
	}
}
