package binaryTree


//疑问： 二叉树节点上的值重复？

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	if len(inorder) == 1 && len(postorder) == 1 {
		root.Val = inorder[0]
		return root
	}

	//Find root node in inorder
	rootindex := 0
	rootValue := postorder[len(postorder)-1]
	for i, v := range inorder {
		if v == rootValue {
			root.Val = v
			rootindex = i
		}
	}

	leftInorder := inorder[:rootindex]
	rightInorder := inorder[rootindex+1:]
	leftPostorder := postorder[:len(leftInorder)]
	rightPostorder := postorder[len(leftInorder) : len(postorder)-1]

	root.Left = buildTree(leftInorder, leftPostorder)
	root.Right = buildTree(rightInorder, rightPostorder)

	return root
}
