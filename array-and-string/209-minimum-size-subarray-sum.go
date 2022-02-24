package main

import (
	"fmt"
)

func main() {
	fmt.Println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))
}

//前缀和+快慢指针
func minSubArrayLen(target int, nums []int) int {
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
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
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
