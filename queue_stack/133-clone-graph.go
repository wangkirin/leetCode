package main

func main() {

}

type Node struct {
	Val       int
	Neighbors []*Node
}

//DFS递归解法
func cloneGraph(node *Node) *Node {
	visited := map[*Node]*Node{}

	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if _, ok := visited[node]; ok {
			return visited[node]
		}
		clone := &Node{node.Val, []*Node{}}
		visited[node] = clone
		for _, n_node := range node.Neighbors {
			visited[node].Neighbors = append(visited[node].Neighbors, dfs(n_node))
		}
		return clone
	}

	return dfs(node)
}

//func cloneGraph(node *Node) *Node {
//	if node == nil {
//		return nil
//	}
//	start := &Node{
//		Val:       node.Val,
//		Neighbors: nil,
//	}
//	if len(node.Neighbors) > 0 {
//		start.Neighbors = make([]*Node, len(node.Neighbors))
//		for i, neighbor := range node.Neighbors {
//			start.Neighbors[i] = cloneGraph(neighbor)
//		}
//	}
//
//	return start
//}

//func cloneGraph(node *Node) *Node {
//	if node == nil {
//		return nil
//	}
//	start := new(Node)
//	start = nodeDfs(node, start)
//	return start
//}
//
//func nodeDfs(node *Node, copyNode *Node) *Node {
//	if node == nil {
//		return nil
//	}
//	copyNode.Val = node.Val
//	if len(node.Neighbors) > 0 {
//		copyNode.Neighbors = make([]*Node, len(node.Neighbors))
//		for i, neighbor := range node.Neighbors {
//			copyNode.Neighbors[i] = nodeDfs(neighbor, copyNode.Neighbors[i])
//		}
//	}
//	return copyNode
//}
