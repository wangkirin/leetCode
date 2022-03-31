package main

import "sort"

func main() {

}

func majorityElement(nums []int) int {
	sort.Ints(nums)
	mid := len(nums) / 2
	if nums[mid] == nums[len(nums)-1] || (nums[mid] != nums[0] && nums[mid] != nums[len(nums)-1]) {
		return nums[mid]
	}
	return nums[0]
}
