package problem0002

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// ListNode2 defines for singly-linked list.
//
//	type ListNode2 struct {
//	    Val int
//	    Next *ListNode2
//	}
type ListNode2 = kit.ListNode

func addTwoNumbersOld(l1 *ListNode2, l2 *ListNode2) *ListNode2 {
	resPre := &ListNode2{}
	cur := resPre
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10

		cur.Next = &ListNode2{Val: sum % 10}
		cur = cur.Next
	}

	return resPre.Next
}
