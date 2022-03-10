package main

import "sort"

func main() {

}

func intersect(nums1 []int, nums2 []int) []int {
	out := []int{}
	sort.Ints(nums1)
	sort.Ints(nums2)
	if len(nums2) < len(nums1) {
		nums1, nums2 = nums2, nums1
	}
	for _, v := range nums1 {
		if searchIII(nums2, v) != -1 {
			out = append(out, v)
		}
	}
	return out
}

func searchIII(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		index := start + (end-start)/2
		if nums[index] == target {
			return index
		}
		if nums[index] > target {
			end = index - 1
		} else {
			start = index + 1
		}
	}
	return -1
}
