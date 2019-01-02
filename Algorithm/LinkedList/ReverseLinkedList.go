package LinkedList

/**
Reverse a singly linked list.

Example:
	Input: 1->2->3->4->5->NULL
	Output: 5->4->3->2->1->NULL

Follow up:
	A linked list can be reversed either iteratively or recursively.
 */

/**
type ListNode struct {
	Val int
	Next *ListNode
}
 */
func reverseList(head *ListNode) *ListNode {
	cur := head
	var result, temp *ListNode
	for cur != nil {
		n := &ListNode{Val:cur.Val, Next:temp}
		if cur.Next == nil {
			result = n
		}
		temp = n
		cur = cur.Next
	}
	return result
}