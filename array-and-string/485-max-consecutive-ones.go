package main

import "fmt"

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
}
func findMaxConsecutiveOnes(nums []int) int {
	slow := 0
	max := 0
	for fast := 0; fast < len(nums); {
		if nums[fast] == 0 {
			slow = fast + 1
		}
		fast++
		temp := fast - slow
		if temp > max {
			max = temp
		}
	}
	return max
}
