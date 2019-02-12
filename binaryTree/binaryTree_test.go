package binaryTree

import "testing"

func TestBuildTree(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	tree := buildTree(inorder, postorder)
	t.Log(tree.Val)
	t.Log(tree.Left.Val)
}

func TestBuildTreeFromPreAndInorder(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	preorder := []int{3, 9, 20, 15, 7}
	tree := buildTreePreAndInorder(preorder, inorder)
	t.Log(tree.Val)
	t.Log(tree.Left.Val)
}

func TestMaxDepth(t *testing.T) {
	bTree := new(TreeNode)
	bTree.Val = 3

	l1Left := new(TreeNode)
	l1Left.Val = 9
	bTree.Left = l1Left
	l1Right := new(TreeNode)
	l1Right.Val = 20
	bTree.Right = l1Right

	l1RightLeft := new(TreeNode)
	l1RightLeft.Val = 15
	l1Right.Left = l1RightLeft
	l1RightRight := new(TreeNode)
	l1RightRight.Val = 7
	l1Right.Right = l1RightRight

	t.Log(maxDepth(bTree))

}

func TestLevelOrderCase1(t *testing.T) {
	bTree := new(TreeNode)
	bTree.Val = 3

	l1Left := new(TreeNode)
	l1Left.Val = 9
	bTree.Left = l1Left
	l1Right := new(TreeNode)
	l1Right.Val = 20
	bTree.Right = l1Right

	l1RightLeft := new(TreeNode)
	l1RightLeft.Val = 15
	l1Right.Left = l1RightLeft
	l1RightRight := new(TreeNode)
	l1RightRight.Val = 7
	l1Right.Right = l1RightRight

	out := levelOrder(bTree)
	t.Log(out)
}

func TestLevelOrderCaseOnlyRoot(t *testing.T) {
	bTree := new(TreeNode)
	bTree.Val = 1

	out := levelOrder(bTree)
	t.Log(out)
}
