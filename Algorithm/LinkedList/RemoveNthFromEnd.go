package LinkedList

/**
Given a linked list, remove the n-th node from the end of list and return its head.

Example:
	Given linked list: 1->2->3->4->5, and n = 2.
	After removing the second node from the end, the linked list becomes 1->2->3->5.

Note:
	Given n will always be valid.
 */

/**
type ListNode struct {
	Val int
	Next *ListNode
}
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := ListNode{}
	dummy.Next = head
	first := &dummy
	second := &dummy

	for i := 1; i < n +1; i++ {
		first = first.Next
	}

	for first.Next != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}
