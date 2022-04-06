package main

func main() {

}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, columns := len(matrix), len(matrix[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, columns)
	}

	var (
		total       = rows * columns
		order       = make([]int, total)
		row, column = 0, 0
		//方向数组
		directions = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
		//判断方向
		directionIndex = 0
	)

	for i := 0; i < total; i++ {
		order[i] = matrix[row][column]
		visited[row][column] = true
		nextRow, nextColumn := row+directions[directionIndex][0], column+directions[directionIndex][1]
		if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= columns || visited[nextRow][nextColumn] {
			//转换方向
			directionIndex = (directionIndex + 1) % 4
		}
		row += directions[directionIndex][0]
		column += directions[directionIndex][1]
	}
	return order
}

//func spiralOrder(matrix [][]int) []int {
//	var spOrder func(mat [][]int, reverse bool) []int
//	spOrder = func(mat [][]int, reverse bool) []int {
//		m := len(mat)
//		n := len(mat[0])
//		ans := []int{}
//		if m == 1 {
//			if reverse {
//				return rever(mat[0])
//			}
//			return mat[0]
//		}
//		if n == 1 {
//			for i := 0; i < m; i++ {
//				ans = append(ans, mat[i][0])
//			}
//			if reverse {
//				return rever(ans)
//			}
//			return ans
//		}
//		if !reverse {
//			ans = append(ans, mat[0]...)
//			for i := 1; i < m; i++ {
//				ans = append(ans, mat[i][n-1])
//			}
//			ans = append(ans, spOrder(mat[1:][:n-1], !reverse)...)
//			return ans
//		} else {
//			ans = append(ans, rever(mat[m-1])...)
//			for i := m - 2; i >= 0; i-- {
//				ans = append(ans, mat[0][i])
//			}
//			ans = append(ans, spOrder(mat[:m-1][:n], !reverse)...)
//			return ans
//		}
//	}
//	return spOrder(matrix, false)
//}
//
//func rever(in []int) []int {
//	l := 0
//	r := len(in) - 1
//	for l < r {
//		in[l], in[r] = in[r], in[l]
//		l++
//		r--
//	}
//	return in
//}
