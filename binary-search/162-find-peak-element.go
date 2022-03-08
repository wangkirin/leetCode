package main

func main() {

}

//非有序数组用二分法

//暴力算法
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return 0
		}
		return 1
	}
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] < nums[i+1] && nums[i+1] > nums[i+2] {
			return i + 1
		}
	}
	if nums[len(nums)-1] > nums[0] {
		return len(nums) - 1
	}
	return 0

}

//方法二：求最大值，最大值即为峰值
