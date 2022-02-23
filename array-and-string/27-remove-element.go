package main

func main() {

}

//快慢指针
func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] == val {
			continue
		}
		nums[slow] = nums[fast]
		slow++
	}
	return slow
}

//左右指针
//func removeElement(nums []int, val int) int {
//	left, right := 0, len(nums)
//	for left < right {
//		if nums[left] == val {
//			nums[left] = nums[right-1]
//			right--
//		} else {
//			left++
//		}
//	}
//	return left
//}
