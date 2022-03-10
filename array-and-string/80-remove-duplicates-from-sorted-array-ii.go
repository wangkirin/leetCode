package main

func main() {

}

//对于此类问题，我们应该进行如下考虑：
//
//由于是保留 k 个相同数字，对于前 k 个数字，我们可以直接保留
//对于后面的任意数字，能够保留的前提是：与当前写入的位置前面的第 k 个元素进行比较，不相同则保留
//

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	slow := 2
	for fast := 2; fast < n; fast++ {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
