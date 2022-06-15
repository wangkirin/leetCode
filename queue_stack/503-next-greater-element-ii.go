package main

func main() {

}

// 利用取模构建循环数组
func nextGreaterElements(nums []int) []int {
	ans := make([]int, len(nums))
	stack := []int{}
	for i := 2*len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i%len(nums)] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			ans[i%len(nums)] = -1
		} else {
			ans[i%len(nums)] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i%len(nums)])
	}
	return ans
}
