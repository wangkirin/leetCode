package binaryTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func preorderTraversal2(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	p := root
	for p != nil || len(stack) > 0 {
		for p != nil {
			result = append(result, p.Val)
			stack = append(stack, p)
			p = p.Left
		}
		p = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		p = p.Right
	}
	return result
}
