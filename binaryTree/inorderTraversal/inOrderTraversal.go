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


//递归算法
var result=make([]int,0)

func inorderTraversal(root *TreeNode) []int   {
	if root==nil {
		return nil
	}
	inorder(root)
	return result
}

func inorder(root *TreeNode) {
	if root.Left!=nil{
		inorder(root.Left)
	}
	if root!=nil{
		result=append(result,root.Val)
	}
	if root.Right!=nil{
		inorder(root.Right)
	}
}






