package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}

func minSubArrayLen(target int, nums []int) int {
	preSum := make([]int, len(nums)+1)
	preSum[0] = 0
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	l, minL := 0, math.MaxInt
	flag := false
	for r := 0; r < len(preSum); r++ {
		for l < r && preSum[r]-preSum[l] >= target {
			flag = true
			l++
			minL = min(minL, r-l+1)
		}
	}
	if !flag {
		minL = 0
	}
	return minL
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//前缀和+快慢指针
func minSubArrayLen1(target int, nums []int) int {
	preSum := make([]int, len(nums)+1)
	preSum[0] = 0
	minLen := 0
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
		if preSum[i+1] >= target {
			if minLen == 0 {
				minLen = i + 1
			}
		}
	}
	for slow := 0; slow < len(nums); slow++ {
		fast := slow + 1
		for fast < len(preSum) {
			if preSum[fast]-preSum[slow] >= target {
				if minLen == 0 {
					minLen = fast - slow
				} else {
					minLen = min(minLen, fast-slow)
				}
				break
			}
			fast++
		}
	}
	return minLen
}

//滑窗
//func minSubArrayLen(s int, nums []int) int {
//	n := len(nums)
//	if n == 0 {
//		return 0
//	}
//	ans := math.MaxInt32
//	start, end := 0, 0
//	sum := 0
//	for end < n {
//		sum += nums[end]
//		for sum >= s {
//			ans = min(ans, end - start + 1)
//			sum -= nums[start]
//			start++
//		}
//		end++
//	}
//	if ans == math.MaxInt32 {
//		return 0
//	}
//	return ans
//}
//
//func min(x, y int) int {
//	if x < y {
//		return x
//	}
//	return y
//}
