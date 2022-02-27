package main

import "fmt"

func main() {
	fmt.Println(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))
	fmt.Println(canVisitAllRooms([][]int{{2}, {}, {1}}))
}

// 深度优先搜索
//我们可以使用深度优先搜索的方式遍历整张图，统计可以到达的节点个数，并利用数组 {vis}vis 标记当前节点是否访问过，以防止重复访问。

var (
	num int
	vis []bool
)

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	num = 0
	vis = make([]bool, n)
	dfs(rooms, 0)
	return num == n
}

func dfs(rooms [][]int, x int) {
	vis[x] = true
	num++
	for _, it := range rooms[x] {
		if !vis[it] {
			dfs(rooms, it)
		}
	}
}

//广度优先搜索
//我们也可以使用广度优先搜索的方式遍历整张图，统计可以到达的节点个数，并利用数组 \textit{vis}vis 标记当前节点是否访问过，以防止重复访问。
func canVisitAllRoomsBFS(rooms [][]int) bool {
	n := len(rooms)
	num := 0
	vis := make([]bool, n)
	queue := []int{}
	vis[0] = true
	queue = append(queue, 0)
	for i := 0; i < len(queue); i++ {
		x := queue[i]
		num++
		for _, it := range rooms[x] {
			if !vis[it] {
				vis[it] = true
				queue = append(queue, it)
			}
		}
	}
	return num == n
}

//思路错误

//func canVisitAllRooms(rooms [][]int) bool {
//	roomNum := len(rooms)
//	matrix := make([][]int, roomNum)
//	for i := range matrix {
//		matrix[i] = make([]int, roomNum)
//	}
//	for i := 0; i < roomNum; i++ {
//		for _, v := range rooms[i] {
//			matrix[i][v] = 1
//		}
//	}
//	fmt.Println(matrix)
//	for i := 1; i < roomNum; i++ {
//		sum := 0
//		for j := 0; j < i; j++ {
//			if matrix[j][i] == 1 {
//				sum += 1
//			}
//		}
//		if sum == 0 {
//			return false
//		}
//	}
//	return true
//}
