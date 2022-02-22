package main

import "fmt"

func main() {
	testcase := []int{1, 7, 3, 6, 5, 6}
	testcase1 := []int{1, 2, 3}
	testcase2 := []int{2, 1, -1}
	fmt.Println(findMiddleIndex(testcase))
	fmt.Println(findMiddleIndex(testcase1))
	fmt.Println(findMiddleIndex(testcase2))
}

func findMiddleIndex(nums []int) int {
	preSum := make([]int, len(nums)+1)
	preSum[0] = 0
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	total := preSum[len(nums)]
	for i, _ := range nums {
		i += 1
		if preSum[i-1] == (total - preSum[i]) {
			return i - 1
		}
	}
	return -1
}
