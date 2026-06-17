package hard

import (
	"fmt"
	"testing"
)

func Test_23(t *testing.T) {
	/*lists := append([]*ListNode{}, GetList([]int{1, 4, 5}))
	lists = append(lists, GetList([]int{1, 3, 4}))
	lists = append(lists, GetList([]int{2, 6}))*/

	lists := append([]*ListNode{}, GetList([]int{3}))
	lists = append(lists, GetList([]int{0}))

	mlist := mergeKLists(lists)
	fmt.Println(mlist)
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}

	var top *ListNode

	for _, list := range lists {
		if list != nil {
			curr := list
			for curr != nil {
				top = insertNode(top, curr.Val)
				curr = curr.Next
			}
		}
	}
	return top
}

func insertNode(top *ListNode, val int) *ListNode {
	if top == nil {
		return &ListNode{
			Val: val,
		}
	}

	curr := top
	var prev *ListNode

	for {
		if curr == nil {
			prev.Next = &ListNode{
				Val: val,
			}
			top = prev.Next //
			break
		} else if curr.Val <= val {
			prev = curr
			curr = curr.Next
		} else {
			if prev != nil {
				prev.Next = &ListNode{
					Val:  val,
					Next: curr,
				}
				top = prev.Next //
			} else {
				top = &ListNode{
					Val:  val,
					Next: curr,
				}
			}
			break
		}
	}
	return top
}
