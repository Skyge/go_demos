package LinkedList

/**
Given a singly linked list, determine if it is a palindrome.

Example 1:

Input: 1->2
Output: false
Example 2:

Input: 1->2->2->1
Output: true
*/
func isPalindrome(head *ListNode) bool {
	if head == nil { // linked list length == 1
		return true
	}
	slow, fast, prev := head, head, &ListNode{}
	prev = nil
	for fast != nil {
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = nil
		}
	}

	if slow == nil { // linked list length == 1
		return true
	}
	for slow != nil {
		slow.Next, prev, slow = prev, slow, slow.Next
	}

	for head != nil && prev != nil {
		if head.Val != prev.Val {
			return false
		}
		head, prev = head.Next, prev.Next
	}
	return true
}
