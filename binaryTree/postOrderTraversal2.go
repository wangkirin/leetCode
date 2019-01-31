package binaryTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 思路：先移动到右结点最底，然后向上（类似前序遍历）,最后以根结点为中心，对称互换

func postorderTraversal2(root *TreeNode) []int {
	var ret []int
	var stack []*TreeNode

	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			ret = append(ret, curr.Val)
			curr = curr.Right
		}

		// pop from stack
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		curr = curr.Left
	}

	return reverse(ret)
}

func reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
