package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(myAtoi("  -42"))
}

//func myAtoi(s string) int {
//	sum, sign, i, n := 0, 1, 0, len(s)
//	// 丢弃无用的前导空格
//	for i < n && s[i] == ' ' {
//		i++
//	}
//	// 判定正负
//	if i < n {
//		if s[i] == '-' {
//			sign = -1
//			i++
//		} else if s[i] == '+' {
//			sign = 1
//			i++
//		}
//	}
//	// 从左到右依次累加
//	for i < n && s[i] >= '0' && s[i] <= '9' {
//		sum = 10*sum + int(s[i]-'0')
//		// 整数超过32位有符号整数范围,特殊处理
//		if sign*sum < math.MinInt32 {
//			return math.MinInt32
//		} else if sign*sum > math.MaxInt32 {
//			return math.MaxInt32
//		}
//		i++
//	}
//	return sign * sum
//}
//
func myAtoi(s string) int {
	strings.TrimPrefix(s, " ")
	out := 0
	nega := false
	ss := []byte(s)
	if len(ss) == 0 {
		return 0
	}
	for i := 0; ss[i] != ' '; i++ {

		if ss[i] == '-' {
			nega = true
		}
		if ss[i] >= '0' && ss[i] <= '9' {
			out = out*10 + int(ss[i]-'0')
		}
		if out < math.MinInt32/10 || out > math.MaxInt32/10 {
			return 0
		}

	}
	if nega {
		out = -1 * out
	}
	return out
}
