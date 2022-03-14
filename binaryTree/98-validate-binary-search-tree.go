package main

func main() {

}

//递归
func isValidBST(root *TreeNode) bool {
	return isValidBST1(root, nil, nil)
}

func isValidBST1(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return isValidBST1(root.Left, min, root) &&
		isValidBST1(root.Right, root, max)
}

//方法二：二叉搜索树「中序遍历」得到的值构成的序列一定是升序的，这启示我们在中序遍历的时候实时检查当前节点的值是否大于前一个中序遍历到的节点的值即可。
