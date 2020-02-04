package queue_stack

import "container/list"

//No. 200
func numIslands(grid [][]byte) int {
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
						if (row+1) <nr && grid[row+1][col] == '1' {
							queue.PushBack((row+1)*nc + col)
							grid[row+1][col] = '0'
						}
						if (col-1) >= 0 && grid[row][col-1] == '1' {
							queue.PushBack(row*nc + col - 1)
							grid[row][col-1] = '0'
						}
						if (col+1) <nc && grid[row][col+1] == '1' {
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
