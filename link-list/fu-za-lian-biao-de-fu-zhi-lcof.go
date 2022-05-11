package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	ans := new(Node)
	for head.Next != nil {
		n := new(Node)
		ans.Val = head.Val
		ans.Next = head.Next
		ans.Random = head.Random
		head = head.Next
		ans = ans.Next
	}

	return ans
}
