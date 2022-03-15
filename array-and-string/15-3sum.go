package main

import (
	"sort"
)

func main() {
	threeSum([]int{-1, 0, 1, 2, -1, -4})
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	out := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		tworesult := twosumTarget(nums[i:], 0-nums[i])
		for _, ints := range tworesult {
			ints = append(ints, nums[i])
			out = append(out, ints)
		}
		//去重
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return out
}
func twosumTarget(nums []int, target int) [][]int {
	out := [][]int{}
	sort.Ints(nums)
	l := 0
	r := len(nums) - 1
	for l < r {
		sum := nums[l] + nums[r]
		left := nums[l]
		right := nums[r]
		if sum == target {
			out = append(out, []int{nums[l], nums[r]})
			//跳过重复元素
			for l < r && nums[l] == left {
				l++
			}
			for l < r && nums[r] == right {
				r--
			}
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return out
}
