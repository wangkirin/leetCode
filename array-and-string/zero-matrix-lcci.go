package main

//https://leetcode-cn.com/problems/zero-matrix-lcci/
func main() {

}
func setZeroes(matrix [][]int) {
	rowZeros := make(map[int]int)
	colZeros := make(map[int]int)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rowZeros[i] = 1
				colZeros[j] = 1
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if rowZeros[i] == 1 || colZeros[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}
