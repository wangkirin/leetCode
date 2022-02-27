package numIslands

import "container/list"

//No. 200

//DFS方法
//我们可以将二维网格看成一个无向图，竖直或水平相邻的 11 之间有边相连。
//
//为了求出岛屿的数量，我们可以扫描整个二维网格。如果一个位置为 11，则以其为起始节点开始进行深度优先搜索。在深度优先搜索的过程中，每个搜索到的 11 都会被重新标记为 00。
//
//最终岛屿的数量就是我们进行深度优先搜索的次数。

func numIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func dfs(grid [][]byte, r, c int) {
	h, w := len(grid), len(grid[0])
	if r < 0 || r >= h || c < 0 || c >= w {
		return
	}
	if grid[r][c] == '0' {
		return
	}
	grid[r][c] = '0'
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}

////BFS方法
func numIslandsBFS(grid [][]byte) int {
	nr := len(grid)
	if nr == 0 {
		return 0
	}
	nc := len(grid[0])
	islandNum := 0

	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				islandNum++
				grid[r][c] = '0'
				queue := list.New()
				queue.PushBack(r*nc + c)
				for {
					if queue.Len() > 0 {
						hash := queue.Front().Value.(int)
						queue.Remove(queue.Front())
						row := hash / nc
						col := hash % nc
						if (row-1) >= 0 && grid[row-1][col] == '1' {
							queue.PushBack((row-1)*nc + col)
							grid[row-1][col] = '0'
						}
						if (row+1) < nr && grid[row+1][col] == '1' {
							queue.PushBack((row+1)*nc + col)
							grid[row+1][col] = '0'
						}
						if (col-1) >= 0 && grid[row][col-1] == '1' {
							queue.PushBack(row*nc + col - 1)
							grid[row][col-1] = '0'
						}
						if (col+1) < nc && grid[row][col+1] == '1' {
							queue.PushBack(row*nc + col + 1)
							grid[row][col+1] = '0'
						}

					} else {
						break
					}
				}

			}
		}
	}
	return islandNum
}
