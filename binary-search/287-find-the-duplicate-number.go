package main

func main() {

}

//二分查找
//二分查找的思路是先猜一个数（有效范围 [left..right] 里位于中间的数 mid），然后统计原始数组中 小于等于 mid 的元素的个数 cnt：
//
//如果 cnt 严格大于 mid。根据抽屉原理，重复元素就在区间 [left..mid] 里；
//否则，重复元素就在区间 [mid + 1..right] 里。

func findDuplicate(nums []int) int {
	n := len(nums)
	l, r := 1, n-1
	ans := -1
	for l <= r {
		mid := (l + r) >> 1
		cnt := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			l = mid + 1
		} else {
			r = mid - 1
			ans = mid
		}
	}
	return ans
}

//方法一：暴力
//func findDuplicate(nums []int) int {
//	for i := 0; i < len(nums)-1; i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[j] == nums[i] {
//				return nums[i]
//			}
//		}
//	}
//	return 0
//}

//方法二:哈希表
