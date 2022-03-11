package main

func main() {

}

//方法一：单指针
//思路与算法
//
//我们可以考虑对数组进行两次遍历。在第一次遍历中，我们将数组中所有的 0 交换到数组的头部。在第二次遍历中，我们将数组中所有的 1 交换到头部的 0 之后。
//此时，所有的 2 都出现在数组的尾部，这样我们就完成了排序。
//func swapColors(colors []int, target int) (countTarget int) {
//	for i, c := range colors {
//		if c == target {
//			colors[i], colors[countTarget] = colors[countTarget], colors[i]
//			countTarget++
//		}
//	}
//	return
//}
//
//func sortColors(nums []int) {
//	count0 := swapColors(nums, 0) // 把 0 排到前面
//	swapColors(nums[count0:], 1)  // nums[:count0] 全部是 0 了，对剩下的 nums[count0:] 把 1 排到前面
//}

//双指针

func sortColors(nums []int) {
	p0, p1 := 0, 0
	for i, c := range nums {
		if c == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if c == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}

//基数排序法
//可以采用基数排序法的思想，用一个数组记录下 0，1，3 的次数，后重排，这个算法对数组进行了两次遍历，其实有一种只进行一次遍历的方法。

func sortColors(nums []int) {
	counts := make([]int, 3)
	for _, num := range nums {
		switch num {
		case 0:
			counts[0]++
		case 1:
			counts[1]++
		case 2:
			counts[2]++
		}
	}
	for i := 0; i < len(nums); i++ {
		if i < counts[0] {
			nums[i] = 0
		}
		if i >= counts[0] && i < (counts[0]+counts[1]) {
			nums[i] = 1
		}
		if i >= (counts[0]+counts[1]) && i < (counts[0]+counts[1]+counts[2]) {
			nums[i] = 2
		}
	}
}
