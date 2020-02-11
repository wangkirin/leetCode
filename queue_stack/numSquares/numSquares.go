package numSquares

import (
	"math"
)

// NO.279
//https://www.cxyxiaowu.com/6940.html
//使用广度优先搜索方法，将 n 依次减去比 n 小的所有平方数，直至 n = 0 ，此时的层数即为最后的结果。

func numSquares(n int) int {
	round := 1
	queue := []int{}
	initMaxSquare := maxSquare(n)
	if initMaxSquare == n {
		return round
	} else {
		queue = append(queue, n)
	}
	for len(queue) > 0 {
		round++
		queuelen := len(queue)
		for i := 0; i < queuelen; i++ {
			square := queue[0]
			maxRoot := maxSquareRoot(square)
			for j := 1; j <= maxRoot; j++ {
				minus := square - j*j
				if minus == maxSquare(minus) {
					return round
				}
				queue = append(queue, minus)
			}
			queue = queue[1:]
		}

	}
	return round
}

func maxSquare(n int) int {
	if n < 0 {
		return 0
	}
	maxSquareRoot := int(math.Floor(math.Sqrt(float64(n))))
	return maxSquareRoot * maxSquareRoot
}
func maxSquareRoot(n int) int {
	if n < 0 {
		return 0
	}
	return int(math.Floor(math.Sqrt(float64(n))))
}

//// 返回 x 的平方根的整数部分
//// 这个函数比 int(math.Sqrt(float64(x))) 快的多
//// 详见 benchmark 的结果
//func intSqrt(x int) int {
//	res := x
//
//	// 牛顿法求平方根
//	for res*res > x {
//		res = (res + x/res) / 2
//	}
//
//	return res
//}