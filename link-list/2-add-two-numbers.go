package main

func main() {

}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	jinwei := 0
	ans := new(ListNode)
	cur := ans
	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		cur.Val = (v1 + v2 + jinwei) % 10
		jinwei = (v1 + v2 + jinwei) / 10

		next := new(ListNode)
		cur.Next = next
		cur = cur.Next
	}
	if jinwei > 0 {
		cur.Next = &ListNode{Val: jinwei}
	}
	return ans
}

//func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
//    var tail *ListNode
//    carry := 0
//    for l1 != nil || l2 != nil {
//        n1, n2 := 0, 0
//        if l1 != nil {
//            n1 = l1.Val
//            l1 = l1.Next
//        }
//        if l2 != nil {
//            n2 = l2.Val
//            l2 = l2.Next
//        }
//        sum := n1 + n2 + carry
//        sum, carry = sum%10, sum/10
//        if head == nil {
//            head = &ListNode{Val: sum}
//            tail = head
//        } else {
//            tail.Next = &ListNode{Val: sum}
//            tail = tail.Next
//        }
//    }
//    if carry > 0 {
//        tail.Next = &ListNode{Val: carry}
//    }
//    return
//}
//
