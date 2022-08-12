package problem0002

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/**
type ListNode struct {
	Val  int
	Next *ListNode
}
*/
type ListNode = kit.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	head = &ListNode{}
	tail = head
	sum, carry := 0, 0
	for l1 != nil || l2 != nil || carry > 0 {
		a, b := 0, 0
		if (l1 != nil) {
			a = l1.Val
			l1 = l1.Next
		}
		if(l2 != nil) {
			b = l2.Val
			l2 = l2.Next
		}
		sum = a + b + carry
		carry = sum / 10
		sum = sum % 10
		tail.Next = &ListNode{Val: sum}
		tail = tail.Next
	}
	return head.Next
}
