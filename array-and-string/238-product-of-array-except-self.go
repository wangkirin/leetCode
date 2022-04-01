package main

func main() {

}

//两次前缀和
func productExceptSelf(nums []int) []int {
	start2end := make([]int, len(nums)+1)
	start2end[0] = 1
	end2start := make([]int, len(nums)+1)
	end2start[len(nums)] = 1
	for i := 0; i < len(nums); i++ {
		start2end[i+1] = start2end[i] * nums[i]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		end2start[i] = end2start[i+1] * nums[i]
	}
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = start2end[i] * end2start[i+1]
	}
	return ans
}

//空间复杂度 O(1) 的方法
//尽管上面的方法已经能够很好的解决这个问题，但是空间复杂度并不为常数。
//由于输出数组不算在空间复杂度内，那么我们可以将 L 或 R 数组用输出数组来计算。先把输出数组当作 L 数组来计算，然后再动态构造 R 数组得到结果
