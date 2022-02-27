package main

func main() {

}

//
//最短路径就用bfs。 先把所有1点值变成-1，后面好用来判断是否来过。。。
//
//把所有0点入队，然后更新0点邻居所有-1点距离为1.
//把1点入队，更新所有邻居-1点距离为2.
//以此类推。。。像石头入水一样波纹散开。。。。

func updateMatrix(mat [][]int) [][]int {
	var next = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	m, n := len(mat), len(mat[0])
	var q [][]int
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 0 {
				q = append(q, []int{i, j})
			} else {
				mat[i][j] -= 2
			}
		}
	}
	//bfs操作
	for len(q) != 0 {
		c := q[0]
		q = q[1:]
		x, y := c[0], c[1]
		for k := range next {
			nx, ny := x+next[k][0], y+next[k][1]
			if nx < 0 || nx == m || ny < 0 || ny == n || mat[nx][ny] != -1 {
				continue
			}
			mat[nx][ny] = mat[x][y] + 1
			q = append(q, []int{nx, ny})
		}
	}
	return mat
}
