package main

import (
	"sort"
)

func main() {
	println(search2([]int{1}, 1))
}
func search2(nums []int, target int) int {
	reverseindex := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			reverseindex = i
			break
		}
	}
	sort.Ints(nums)
	index := searchB(nums, target)
	if index == -1 {
		return -1
	}
	if index < (len(nums) - reverseindex - 1) {
		return index + reverseindex
	}
	return index - (len(nums) - reverseindex)
}

func searchB(nums []int, target int) int {
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

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] {
			if nums[mid] > target && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
