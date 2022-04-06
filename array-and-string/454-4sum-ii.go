package main

func main() {

}

//
//方法一：分组 + 哈希表
//思路与算法
//
//我们可以将四个数组分成两部分，AA 和 BB 为一组，CC 和 DD 为另外一组
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count12 := make(map[int]int)
	for _, a := range nums1 {
		for _, b := range nums2 {
			count12[a+b]++
		}
	}
	ans := 0
	for _, c := range nums3 {
		for _, d := range nums4 {
			ans += count12[-(c + d)]
		}
	}
	return ans
}
