package Trees

/**
Given a binary tree, return the level order traversal of its nodes' values.
(ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
]
*/
func levelOrder(root *TreeNode) [][]int {
	return helper(root)
}

func helper(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int
	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		lenQ := len(q)
		var tmp []int
		for i := 0; i < lenQ; i++ {
			n := q[i]
			tmp = append(tmp, n.Val)
			if n.Left != nil {
				q = append(q, n.Left)
			}

			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		q = q[lenQ:]
		result = append(result, tmp)
	}

	return result
}
