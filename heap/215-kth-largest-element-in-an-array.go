package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

//排序
//func findKthLargest(nums []int, k int) int {
//	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
//	return nums[k-1]
//
//}

////堆
type HeapArray []int

func findKthLargest(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	ha := &HeapArray{}
	heap.Init(ha)
	for _, num := range nums {
		heap.Push(ha, num)

	}
	ans := 0
	for i := 0; i < k; i++ {
		if i == k-1 {
			ans = heap.Pop(ha).(int)
			return ans
		}
		heap.Pop(ha)
	}
	return ans
}
func (h HeapArray) Len() int { return len(h) }
func (h HeapArray) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h HeapArray) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HeapArray) Push(x interface{}) {
	// Push 和 Pop 使用 pointer receiver 作为参数，
	// 因为它们不仅会对切片的内容进行调整，还会修改切片的长度。
	*h = append(*h, x.(int))
}

func (h *HeapArray) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
