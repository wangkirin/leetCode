package main

import "sort"

func main() {

}

//二分查找
func intersection(nums1 []int, nums2 []int) []int {
	out := []int{}
	sort.Ints(nums1)
	sort.Ints(nums2)
	last := -1
	for _, v := range nums1 {
		if last == v {
			continue
		}
		last = v
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

// 两个哈希表
//func intersection(nums1 []int, nums2 []int) (intersection []int) {
//	set1 := map[int]struct{}{}
//	for _, v := range nums1 {
//		set1[v] = struct{}{}
//	}
//	set2 := map[int]struct{}{}
//	for _, v := range nums2 {
//		set2[v] = struct{}{}
//	}
//	if len(set1) > len(set2) {
//		set1, set2 = set2, set1
//	}
//	for v := range set1 {
//		if _, has := set2[v]; has {
//			intersection = append(intersection, v)
//		}
//	}
//	return
//}
