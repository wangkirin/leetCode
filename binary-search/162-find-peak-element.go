package main

import (
	"math"
	"math/rand"
)

func main() {

}

//迭代爬坡用二分法优化
func findPeakElement(nums []int) int {
	n := len(nums)
	idx := rand.Intn(n)

	// 辅助函数，输入下标 i，返回 nums[i] 的值
	// 方便处理 nums[-1] 以及 nums[n] 的边界情况
	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}

	for !(get(idx-1) < get(idx) && get(idx) > get(idx+1)) {
		if get(idx) < get(idx+1) {
			idx++
		} else {
			idx--
		}
	}

	return idx
}

//寻找最大值
//func findPeakElement(nums []int) (idx int) {
//	for i, v := range nums {
//		if v > nums[idx] {
//			idx = i
//		}
//	}
//	return
//}

//暴力算法
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return 0
		}
		return 1
	}
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] < nums[i+1] && nums[i+1] > nums[i+2] {
			return i + 1
		}
	}
	if nums[len(nums)-1] > nums[0] {
		return len(nums) - 1
	}
	return 0

}

//方法二：求最大值，最大值即为峰值
