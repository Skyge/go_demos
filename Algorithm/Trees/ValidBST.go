package Trees

/**
Given a binary tree, determine if it is a valid binary search tree (BST).
Assume a BST is defined as follows:
    The left subtree of a node contains only nodes with keys less than the node's key.
    The right subtree of a node contains only nodes with keys greater than the node's key.
    Both the left and right subtrees must also be binary search trees.
Example 1:
	Input:
    	 2
		/ \
	   1   3
Output: true

Example 2:

    5
   / \
  1   4
     / \
    3   6
Output: false
Explanation: The input is: [5,1,4,null,null,3,6]. The root node's value
             is 5 but its right child's value is 4.
*/
func isValidBST(root *TreeNode) bool {
	platformType := 32 << (^uint(0) >> 63)
	var intMin = 0
	var intMax = 0
	if platformType == 64 {
		intMin = -9223372036854775808
		intMax = 9223372036854775807
	} else if platformType == 32 {
		intMin = -2147483648
		intMax = 2147483647
	}

	return helper(root, intMin, intMax)
}

func helper(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val < min || root.Val > max {
		return false
	}
	return helper(root.Left, min, root.Val-1) && helper(root.Right, root.Val+1, max)
}
