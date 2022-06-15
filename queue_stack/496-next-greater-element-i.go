package main

func main() {

}

// 单调栈

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := []int{}
	m := make(map[int]int)
	stack := []int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] < nums2[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			m[nums2[i]] = -1
		} else {
			m[nums2[i]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	for _, i := range nums1 {
		ans = append(ans, m[i])
	}
	return ans
}

// 暴力解法1优化
func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		ans[i] = -1
		f := false
		for j := 0; j < len(nums2); j++ {
			if nums2[j] == nums1[i] {
				f = true
				continue
			}
			if f && nums1[i] < nums2[j] {
				ans[i] = nums2[j]
				break
			}
		}
	}
	return ans
}

// 暴力解法1
func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	ans := []int{}
	for i := 0; i < len(nums1); i++ {
		f := false
		for j := 0; j < len(nums2); j++ {
			if nums2[j] == nums1[i] {
				f = true
				continue
			}
			if f && nums1[i] < nums2[j] {
				ans = append(ans, nums2[j])
				f = false
				break
			}
		}
		if f {
			ans = append(ans, -1)
		}
	}
	return ans
}
