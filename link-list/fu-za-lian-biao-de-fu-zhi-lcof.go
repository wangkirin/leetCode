package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//拆分+拼接
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		n := &Node{Val: cur.Val}
		n.Next = cur.Next
		cur.Next = n
		cur = cur.Next.Next
	}
	cur = head
	for cur != nil {
		if cur.Random == nil {
			cur.Next.Random = nil
		} else {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	newHead := head.Next
	for node := head; node != nil; node = node.Next {
		nodeNew := node.Next
		node.Next = node.Next.Next
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next
		}
	}
	return newHead
}

//哈希
func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodemap := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		n := new(Node)
		n.Val = cur.Val
		nodemap[cur] = n
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		nodemap[cur].Next = nodemap[cur.Next]
		nodemap[cur].Random = nodemap[cur.Random]
		cur = cur.Next
	}

	return nodemap[head]
}
