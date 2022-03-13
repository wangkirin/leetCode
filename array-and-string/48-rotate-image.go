package main

func main() {

}

func rotate(matrix [][]int) {
	n := len(matrix)
	// 水平翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

//两轮旋转：第一轮围绕 “左下-右上”轴对掉
//第二轮围绕水平中轴对调
//
//func rotate(matrix [][]int) {
//	n := len(matrix)
//	//	第一轮
//	for i := 0; i < n-1; i++ {
//		for j := 0; j < n-i-1; j++ {
//			matrix[i][j], matrix[n-j][n-i] =  matrix[n-j][n-i], matrix[i][j]
//		}
//	}
//	//	第二轮
//	for i := 0; i < n/2; i++ {
//			matrix[i], matrix[n-1-i]=matrix[n-1-i],matrix[i]
//		}
//	}
//}
