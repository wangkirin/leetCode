package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findClosestElements([]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5))
}

//二分查找
//1.只要找到左边界，加上k，就可以得到右边界。
//2. 使用二分法，对于当前的mid，对比mid+k到x的距离，如果
//	1）x- arr[mid] > arr[mid+k]-x，说明左边界可以右移，left = mid+1
//  2）x- arr[mid] <= arr[mid+k]-x，则反之，right = mid
//3. 截取arr从left到left + k部分

func findClosestElements(arr []int, k int, x int) []int {
	l := len(arr)
	left, right := 0, l-k
	for left < right {
		mid := (left + right) / 2
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return arr[left : left+k]
}

//二维数组，排序,sort函数
func findClosestElements(arr []int, k int, x int) []int {
	sort.Slice(arr, func(i, j int) bool {
		if abs(arr[i]-x) == abs(arr[j]-x) {
			return arr[i] < arr[j]
		}
		return abs(arr[i]-x) < abs(arr[j]-x)
	})
	out := arr[:k]
	sort.Ints(out)
	return out
}

//func findClosestElements(arr []int, k int, x int) []int {
//	a := make([][2]int, len(arr))
//	for i := 0; i < len(arr); i++ {
//		a[i] = [2]int{abs(arr[i] - x), i}
//	}
//	sort.Slice(a, func(i, j int) bool {
//		if a[i][0] == a[j][0] {
//			return a[i][1] < a[j][1]
//		}
//		return a[i][0] < a[j][0]
//	})
//	fmt.Println(a)
//	out := make([]int, k)
//	for j := 0; j < k; j++ {
//		out[j] = arr[a[j][1]]
//	}
//	sort.Ints(out)
//	return out
//}
func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
