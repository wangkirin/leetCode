package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

//排序
func findKthLargest(nums []int, k int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return nums[k-1]

}

////堆
//func findKthLargest(nums []int, k int) int {
//	d := &heapInt2{}
//	for _, v := range nums {
//		heap.Push(d, v)
//	}
//	for i := 0; i < k-1; i++ {
//		heap.Pop(d)
//	}
//	return d.Peek()
//}
//
//type heapInt2 []int
//
////Less  小于就是小跟堆，大于号就是大根堆
//func (h *heapInt2) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
//func (h *heapInt2) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
//func (h *heapInt2) Len() int           { return len(*h) }
//func (h *heapInt2) Push(x interface{}) {
//	*h = append(*h, x.(int))
//}
//func (h *heapInt2) Pop() interface{} {
//	t := (*h)[len(*h)-1]
//	*h = (*h)[:len(*h)-1]
//	return t
//}
//func (h *heapInt2) Peek() int {
//	return (*h)[0]
//}
