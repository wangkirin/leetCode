package main

import "sort"

func main() {

}

//单调栈
func mostCompetitive(nums []int, k int) []int {
	stack := []int{}
	n := len(nums)
	last := n - k
	for i := 0; i < n; i++ {
		for len(stack) > 0 && last > 0 && stack[len(stack)-1] > nums[i] {
			stack = stack[:len(stack)-1]
			last--
		}
		stack = append(stack, nums[i])
	}
	return stack[:k]
}

//会超时 85/88
func mostCompetitive1(nums []int, k int) []int {
	ans := []int{}
	var dfs func(subnums []int, l int)
	dfs = func(subnums []int, l int) {
		if l == 1 {
			sort.Ints(subnums)
			ans = append(ans, subnums[0])
			return
		}
		n := len(subnums)
		minindex := 0
		for i := 0; i < n-l+1; i++ {
			if subnums[i] < subnums[minindex] {
				minindex = i

			}
		}
		ans = append(ans, subnums[minindex])
		dfs(subnums[minindex+1:], l-1)
	}
	dfs(nums, k)
	return ans
}
