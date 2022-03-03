package main

import "sort"

func main() {

}

//哈希表
func containsDuplicate(nums []int) bool {
	nummap := make(map[int]int)
	for _, v := range nums {
		if nummap[v] == 1 {
			return true
		}
		nummap[v] = 1
	}
	return false
}

//排序，连续两个位置相等返回
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	if len(nums) == 1 {
		return false
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
