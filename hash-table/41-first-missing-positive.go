package main

import "math"

func main() {

}
func firstMissingPositive(nums []int) int {
	maxnum := math.MaxInt64
	minnum := math.MinInt64
	for _, num := range nums {
		if num >= maxnum {
			maxnum = num
		}
		if num <= minnum {
			minnum = num
		}
	}
	if minnum > 1 {
		return 1
	}
	if maxnum < 0 {

	}
}
