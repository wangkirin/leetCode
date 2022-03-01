package binaryTree

//func maxDepth(root *TreeNode) int {
//	sum := 0
//	var maxfunc func(rootNode *TreeNode) int
//	maxfunc = func(rootNode *TreeNode) int {
//		if rootNode.Left == nil && rootNode.Right == nil {
//			return 1
//		}
//		return max(maxfunc(rootNode.Left), maxfunc(rootNode.Right))
//
//	}
//	return sum + maxfunc(root)
//}
//func max(x, y int) int {
//	if x >= y {
//		return x
//	}
//	return y
//}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	var depthLeft, depthRight int
	if root.Left != nil {
		depthLeft = maxDepth(root.Left)
		depthLeft++
	}
	if root.Right != nil {
		depthRight = maxDepth(root.Right)
		depthRight++
	}
	depth := depthLeft
	if depthLeft < depthRight {
		depth = depthRight
	}
	return depth
}
