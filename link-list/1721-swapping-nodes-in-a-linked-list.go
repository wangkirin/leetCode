package main

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//快慢指针：快指针先行k步,直到k
func swapNodes(head *ListNode, k int) *ListNode {
	fast := head
	for i := 1; i < k; i++ {
		fast = fast.Next
	}
	temp := fast
	slow := head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	temp.Val, slow.Val = slow.Val, temp.Val
	return head
}

//暴力 三次遍历
func swapNodes1(head *ListNode, k int) *ListNode {
	headCopy := head
	count := 1
	cur := head
	for cur.Next != nil {
		count++
		cur = cur.Next
	}
	valA, valB := 0, 0
	cur = head
	for i := 0; i < count && cur != nil; i++ {
		if i == k-1 {
			valA = cur.Val
		}
		if i == count-k {
			valB = cur.Val
		}
		cur = cur.Next
	}
	cur = head
	for i := 0; i < count && cur != nil; i++ {
		if i == k-1 {
			cur.Val = valB
		}
		if i == count-k {
			cur.Val = valA
		}
		cur = cur.Next
	}
	return headCopy
}
