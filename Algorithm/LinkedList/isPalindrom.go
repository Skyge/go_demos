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
	if head == nil || head.Next == nil {
		return true
	}
	if head.Next.Next == nil {
		return head.Val == head.Next.Val
	}

	fast := head
	slow := head
	tag := slow.Next

	for fast.Next != nil || fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast.Next == nil {
		slow = slow.Next
	}
	fast = tag
	for slow != nil {
		if slow.Val == fast.Val {
			return false
		}else {
			slow = slow.Next
			fast = fast.Next
		}
	}
	return true
}