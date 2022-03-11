package main

func main() {

}

//双指针算法
func twoSum(numbers []int, target int) []int {
	slow := 0
	fast := len(numbers) - 1
	for slow < fast && fast < len(numbers) {
		if numbers[slow]+numbers[fast] == target {
			return []int{slow + 1, fast + 1}
		} else if numbers[slow]+numbers[fast] > target {
			fast--
		} else {
			slow++
		}
	}
	return nil
}

//二分查找
//遍历每个 nums[i]，在剩余数组中查找 target-nums[i] 的值，时间复杂度为 O(n log n)
//func twoSum(numbers []int, target int) []int {
//	for i := 0; i < len(numbers); i++ {
//		low, high := i + 1, len(numbers) - 1
//		for low <= high {
//			mid := (high - low) / 2 + low
//			if numbers[mid] == target - numbers[i] {
//				return []int{i + 1, mid + 1}
//			} else if numbers[mid] > target - numbers[i] {
//				high = mid - 1
//			} else {
//				low = mid + 1
//			}
//		}
//	}
//	return []int{-1, -1}
//}
