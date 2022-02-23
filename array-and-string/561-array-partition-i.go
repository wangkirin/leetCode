package main

import "sort"

func main() {

}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

//排序，取偶数
//func arrayPairSum(nums []int) int {
//	sort.Ints(nums)
//	sum := 0
//	for i := 0; i < len(nums); i++ {
//		if i%2 == 0 {
//			sum += nums[i]
//		}
//	}
//	return sum
//}
