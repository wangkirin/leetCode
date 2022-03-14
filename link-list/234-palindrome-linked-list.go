package main

func main() {

}

func isPalindrome(head *ListNode) bool {
	r := reverse(head)
	for head != nil && head.Next != nil {
		if r.Val != head.Val {
			return false
		}
		r.Next.Next = r.Next
		head.Next.Next = head.Next
	}
	return true
}

//方法一：将值复制到数组中后用双指针法
