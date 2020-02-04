package binaryTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Reference: https://blog.csdn.net/zhangxiangDavaid/article/details/37115355
// Memory Limit Exceeded

func postorderTraversal(root *TreeNode) []int {

	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	p := root
	var lastVisit *TreeNode
	lastVisit = nil
	for p != nil || len(stack) > 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left

		}
		p = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		if p.Right == nil || p.Right == lastVisit {
			result = append(result, p.Val)
			lastVisit = p
		} else {
			stack = append(stack, p)
			p = p.Right
			for p != nil {
				stack = append(stack, p)
				p = p.Left
			}
		}
	}
	return result
}
