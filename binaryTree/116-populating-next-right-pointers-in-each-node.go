package main

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//二叉树的层次遍历做法
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	// 初始化队列同时将第一层节点加入队列中，即根节点
	queue := []*Node{root}

	// 循环迭代的是层数
	for len(queue) > 0 {
		tmp := queue
		queue = nil

		// 遍历这一层的所有节点
		for i, node := range tmp {
			// 连接
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}

			// 拓展下一层节点
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	// 返回根节点
	return root
}

//利用next指针
func connect(root *Node) *Node {
	var head = root
	var nextHead = new(Node) //每层链表的哑头结点
	for head != nil {
		var q = nextHead //用于建立下一层链表
		for p := head; p != nil; p = p.Next {
			//遍历本层链表
			if p.Left != nil {
				//非空结点插入下层链表
				q.Next = p.Left
				q = p.Left
			}
			if p.Right != nil {
				//非空结点插入下层链表
				q.Next = p.Right
				q = p.Right
			}
		}
		//本层结束 转移到下一层
		head = nextHead.Next
		//哑头结点复用
		nextHead.Next = nil
	}
	return root
}

