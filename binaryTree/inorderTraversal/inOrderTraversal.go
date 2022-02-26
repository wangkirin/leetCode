package inorderTraversal

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//N0.94

// 迭代算法
//func inorderTraversal(root *TreeNode) []int {
//	result := make([]int, 0)
//	stack := make([]*TreeNode, 0)
//	p := root
//	for p != nil || len(stack) > 0 {
//		for p != nil {
//			stack = append(stack, p)
//			p = p.Left
//		}
//
//		p = stack[len(stack)-1]
//		stack = stack[0 : len(stack)-1]
//		result = append(result, p.Val)
//		p = p.Right
//	}
//	return result
//}

//递归算法（DFS）

func inorderTraversal(root *TreeNode) []int {

	result := []int{}
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		result = append(result, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return result
}
