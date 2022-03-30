package main

func main() {

}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[m][n] == target {
				return true
			}
		}
	}
	return false
}

//Z字形查找
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	x, y := 0, n-1
	for x < m && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}

//func searchMatrix(matrix [][]int, target int) bool {
//	l := [2]int{0, 0}
//	r := [2]int{len(matrix), len(matrix[0])}
//	for l[0] < r[0] && l[1] < r[1] {
//		mid := [2]int{l[0] + (r[0]-l[0])/2, l[1] + (r[1]-l[1])/2}
//		if matrix[mid[0]][mid[1]] == target {
//			return true
//		} else if matrix[mid[0]][mid[1]] > target {
//			return searchMatrix(matrix[])
//		} else if matrix[mid[0]][mid[1]] < target {
//
//		}
//	}
//
//	return false
//}
