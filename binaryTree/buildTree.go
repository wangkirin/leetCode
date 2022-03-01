package binaryTree

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	if len(inorder) == 1 && len(postorder) == 1 {
		root.Val = inorder[0]
		return root
	}

	rootV := postorder[len(postorder)-1]

	root.Val = rootV
	rootindex := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootV {
			rootindex = i
		}
	}
	root.Left = buildTree(inorder[0:rootindex], postorder[0:rootindex])
	root.Right = buildTree(inorder[rootindex+1:], postorder[rootindex:len(postorder)-1])
	return root
}

//疑问： 二叉树节点上的值重复？

//func buildTree(inorder []int, postorder []int) *TreeNode {
//	if len(inorder) == 0 || len(postorder) == 0 {
//		return nil
//	}
//	root := new(TreeNode)
//	if len(inorder) == 1 && len(postorder) == 1 {
//		root.Val = inorder[0]
//		return root
//	}
//
//	//Find root node in inorder
//	rootindex := 0
//	rootValue := postorder[len(postorder)-1]
//	for i, v := range inorder {
//		if v == rootValue {
//			root.Val = v
//			rootindex = i
//		}
//	}
//
//	leftInorder := inorder[:rootindex]
//	rightInorder := inorder[rootindex+1:]
//	leftPostorder := postorder[:len(leftInorder)]
//	rightPostorder := postorder[len(leftInorder) : len(postorder)-1]
//
//	root.Left = buildTree(leftInorder, leftPostorder)
//	root.Right = buildTree(rightInorder, rightPostorder)
//
//	return root
//}

func buildTreePreAndInorder(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 || len(preorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	if len(inorder) == 1 && len(preorder) == 1 {
		root.Val = inorder[0]
		return root
	}

	//Find root node in inorder
	rootindex := 0
	rootValue := preorder[0]
	for i, v := range inorder {
		if v == rootValue {
			root.Val = v
			rootindex = i
		}
	}

	leftInorder := inorder[:rootindex]
	rightInorder := inorder[rootindex+1:]
	leftPreorder := preorder[1 : len(leftInorder)+1]
	rightPreorder := preorder[len(leftInorder)+1 : len(preorder)]

	root.Left = buildTreePreAndInorder(leftPreorder, leftInorder)
	root.Right = buildTreePreAndInorder(rightPreorder, rightInorder)

	return root
}
