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
	// base case
	if head == nil || head.Next == nil {
		return true
	}

	// 1. find the half
	var prev *ListNode = nil
	var slow = head
	var fast = head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	// cut half
	prev.Next = nil

	// 2. set the second half and reverse the first half
	if fast != nil {
		slow = slow.Next
	}
	var head2 *ListNode = slow
	head = reverse(head)

	// 3. compare two linked lists
	for head != nil && head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head = head.Next
		head2 = head2.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil
	var curr = head

	for curr != nil {
		var post = curr.Next
		curr.Next = prev
		prev = curr
		curr = post
	}

	return prev
}
