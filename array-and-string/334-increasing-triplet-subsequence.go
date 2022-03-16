package main

import "math"

func main() {

}

//双向遍历 时间ON，空间ON，一个存储左边最小的，一个存储右边最大的
func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	leftMin := make([]int, n)
	leftMin[0] = nums[0]
	for i := 1; i < n; i++ {
		leftMin[i] = min(leftMin[i-1], nums[i])
	}
	rightMax := make([]int, n)
	rightMax[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], nums[i])
	}
	for i := 1; i < n-1; i++ {
		if nums[i] > leftMin[i-1] && nums[i] < rightMax[i+1] {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

//贪心算法
//从左到右遍历数组nums，遍历过程中维护两个变量first 和second，分别表示递增的三元子序列中的第一个数和第二个数，任何时候都有 first<second。

func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	first, second := nums[0], math.MaxInt32
	for i := 1; i < n; i++ {
		num := nums[i]
		if num > second {
			return true
		} else if num > first {
			second = num
		} else {
			first = num
		}
	}
	return false
}
