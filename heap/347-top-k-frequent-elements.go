package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1}, 1))
}

//func topKFrequent(nums []int, k int) []int {
//	freq := make(map[int]int)
//	for _, num := range nums {
//		freq[num]++
//	}
//	freqints := [][]int{}
//	for k, v := range freq {
//		freqints = append(freqints, []int{k, v})
//	}
//	sort.Slice(freqints, func(i, j int) bool {
//		return freqints[i][1] > freqints[j][1]
//	})
//	out := make([]int, k)
//	for i := 0; i < k; i++ {
//		out[i] = freqints[i][0]
//	}
//	return out
//}

type FreqTimeHeap [][2]int

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}
	h := &FreqTimeHeap{}
	heap.Init(h)
	for k, v := range freqMap {
		heap.Push(h, [2]int{k, v})
	}
	ans := []int{}
	for i := 0; i < k; i++ {
		temp := heap.Pop(h).([2]int)
		ans = append(ans, temp[0])
	}
	return ans
}

func (h FreqTimeHeap) Len() int { return len(h) }
func (h FreqTimeHeap) Less(i, j int) bool {
	return h[i][1] > h[j][1]
}
func (h FreqTimeHeap) Swap(i, j int) {
	h[i][0], h[j][0] = h[j][0], h[i][0]
	h[i][1], h[j][1] = h[j][1], h[i][1]
}

func (h *FreqTimeHeap) Push(x interface{}) {
	// Push 和 Pop 使用 pointer receiver 作为参数，
	// 因为它们不仅会对切片的内容进行调整，还会修改切片的长度。
	*h = append(*h, x.([2]int))
}

func (h *FreqTimeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//堆
//func topKFrequent(nums []int, k int) []int {
//	occurrences := map[int]int{}
//	for _, num := range nums {
//		occurrences[num]++
//	}
//	h := &IHeap{}
//	heap.Init(h)
//	for key, value := range occurrences {
//		heap.Push(h, [2]int{key, value})
//		if h.Len() > k {
//			heap.Pop(h)
//		}
//	}
//	ret := make([]int, k)
//	for i := 0; i < k; i++ {
//		ret[k-i-1] = heap.Pop(h).([2]int)[0]
//	}
//	return ret
//}
//
//type IHeap [][2]int
//
//func (h IHeap) Len() int           { return len(h) }
//func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
//func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//
//func (h *IHeap) Push(x interface{}) {
//	*h = append(*h, x.([2]int))
//}
//
//func (h *IHeap) Pop() interface{} {
//	old := *h
//	n := len(old)
//	x := old[n-1]
//	*h = old[0 : n-1]
//	return x
//}
