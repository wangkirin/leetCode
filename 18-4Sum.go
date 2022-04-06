package main

import "sort"

func main() {

}
func fourSum(nums []int, target int) [][]int {
	return nSum(nums, target, 4)
}

func nSum(nums []int, target int, n int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	if n < 2 && len(nums) < n {
		return ans
	}
	if n == 2 {
		l := 0
		r := len(nums) - 1
		for l < r {
			left := nums[l]
			right := nums[r]
			if nums[l]+nums[r] == target {

				ans = append(ans, []int{nums[l], nums[r]})
				for l < r && nums[l] == left {
					l++
				}
				for l < r && nums[r] == right {
					r--
				}
			} else if nums[l]+nums[r] > target {
				r--

			} else if nums[l]+nums[r] < target {
				l++
			}

		}
	} else {
		for i := 0; i < len(nums); i++ {
			newnums := []int{}
			newnums = append(newnums, nums[i+1:]...)
			results := nSum(newnums, target-nums[i], n-1)
			for _, result := range results {
				result = append(result, nums[i])
				ans = append(ans, result)
			}
			for i < len(nums)-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return ans
}
