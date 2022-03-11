package main

import "sort"

func main() {

}

//逆向双指针
//因此可以指针设置为从后向前遍历，每次取两者之中的较大者放进 nums1的最后面

//双指针
//两个数组看作队列，每次从两个数组头部取出比较小的数字放到结果中。
func merge(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1, sorted)
}

//直接合并后排序
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < len(nums1); i++ {
		if i >= m {
			nums1[i] = nums2[i-m]
		}
	}
	sort.Ints(nums1)
}
