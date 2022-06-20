package main

func main() {

}

//滑窗+前缀和
func longestOnes(nums []int, k int) (ans int) {
	var left, lsum, rsum int
	for right, v := range nums {
		rsum += 1 - v
		for lsum < rsum-k {
			lsum += 1 - nums[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
