package main

func main() {

}

//方法二：多次翻转
func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

//方法一：新数组
func rotate(nums []int, k int) {
	new := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		new[i] = nums[(i+k)%len(nums)]
	}
	copy(nums, new)
}
