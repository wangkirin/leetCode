package binaryTree

import "testing"

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

func TestLevelOrderCase2(t *testing.T) {
	bTree := new(TreeNode)
	bTree.Val = 1

	out := levelOrder(bTree)
	t.Log(out)
}
