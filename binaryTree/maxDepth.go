package binaryTree

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
