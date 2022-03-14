package main

func main() {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	reverse(head)
	temp := head
	for i := 0; i < n; i++ {
		temp = temp.Next
	}
	temp.Val = temp.Next.Val
	temp.Next = temp.Next.Next
	reverse(head)
	return head
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

//方法二：栈
//我们也可以在遍历链表的同时将所有节点依次入栈。根据栈「先进后出」的原则，我们弹出栈的第 nn 个节点就是需要删除的节点，并且目前栈顶的节点就是待删除节点的前驱节点。
//这样一来，删除操作就变得十分方便了。

// 法一：计算链表长度
//func getLength(head *ListNode) (length int) {
//	for ; head != nil; head = head.Next {
//		length++
//	}
//	return
//}
//
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	length := getLength(head)
//	dummy := &ListNode{0, head}
//	cur := dummy
//	for i := 0; i < length-n; i++ {
//		cur = cur.Next
//	}
//	cur.Next = cur.Next.Next
//	return dummy.Next
//}
