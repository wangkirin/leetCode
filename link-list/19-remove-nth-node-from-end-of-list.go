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
