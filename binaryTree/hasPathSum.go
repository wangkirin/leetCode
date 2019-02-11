package binaryTree


//https://leetcode.com/explore/learn/card/data-structure-tree/17/solve-problems-recursively/537/

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			return true
		} else {
			return false
		}
	}
	if hasPathSum(root.Left, (sum-root.Val)) || hasPathSum(root.Right, (sum-root.Val)) {
		return true
	}
	return false
}
