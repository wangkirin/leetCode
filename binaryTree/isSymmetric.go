package binaryTree

//判断一个二叉树是否是沿着根结点对称的，一种方法是分别对根结点的左右子树遍历，判断对称的位置上的值是否相等。而二叉树的遍历，可是分为递归遍历和非递归遍历。
//
//S1:递归

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return SymmetricCompare(root.Left, root.Right)
}

func SymmetricCompare(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left != nil && right != nil {
		return (left.Val == right.Val) && SymmetricCompare(left.Left, right.Right) && SymmetricCompare(left.Right, right.Left)
	}

	return false
}
