package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	testcase := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(testcase, 2))
}
