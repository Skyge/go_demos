package Trees

/**
Given an array where elements are sorted in ascending order, convert it to a height balanced BST.

For this problem, a height-balanced binary tree is defined as a binary tree 
in which the depth of the two subtrees of every node never differ by more than 1.

Example:
    Given the sorted array: [-10,-3,0,5,9],
    One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:

      0
     / \
   -3   9
   /   /
 -10  5
 */

func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    
    root := new(TreeNode)
    root.Val = nums[len(nums)/2]
    root.Left = sortedArrayToBST(nums[:len(nums)/2])
    root.Right = sortedArrayToBST(nums[len(nums)/2+1:])
    return root
}
