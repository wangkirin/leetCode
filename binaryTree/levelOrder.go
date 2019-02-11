package binaryTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	thisLevel := []*TreeNode{}
	nextlevel := []*TreeNode{}
	thisLevel = append(thisLevel, root)
	if root == nil {
		return nil
	}
	if root.Left != nil {
		nextlevel = append(nextlevel, root.Left)
	}
	if root.Right != nil {
		nextlevel = append(nextlevel, root.Right)
	}

	if len(nextlevel) == 0 {
		return [][]int{{root.Val}}
	}

	for len(thisLevel) != 0 && len(nextlevel) != 0 {
		thisLevelOuput := []int{}
		for _, this := range thisLevel {
			thisLevelOuput = append(thisLevelOuput, this.Val)
		}
		result = append(result, thisLevelOuput)

		thisLevel = nextlevel
		nextlevel = nil

		for _, this := range thisLevel {
			if this.Left != nil {
				nextlevel = append(nextlevel, this.Left)
			}
			if this.Right != nil {
				nextlevel = append(nextlevel, this.Right)
			}
		}
		if nextlevel == nil && thisLevel != nil {
			finalLevelOuput := []int{}
			for _, this := range thisLevel {
				finalLevelOuput = append(finalLevelOuput, this.Val)
			}
			result = append(result, finalLevelOuput)
		}
	}
	return result
}

//参考实现，使用一个TreeNode数组

//func levelOrder(root *TreeNode) [][]int {
//	if root == nil {
//		return nil
//	}
//
//	var ret [][]int
//	var q []*TreeNode
//	q = append(q, root)
//	for len(q) > 0 {
//		lenQ := len(q)
//		var tmp []int
//		for i := 0; i < lenQ; i++ {
//			n := q[i]
//			tmp = append(tmp, n.Val)
//			if n.Left != nil {
//				q = append(q, n.Left)
//			}
//
//			if n.Right != nil {
//				q = append(q, n.Right)
//			}
//		}
//		q = q[lenQ:]
//		ret = append(ret, tmp)
//	}
//
//	return ret
//}
