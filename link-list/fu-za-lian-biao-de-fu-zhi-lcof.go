package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
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
