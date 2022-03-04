package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1}, 1))
}

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	freqints := [][]int{}
	for k, v := range freq {
		freqints = append(freqints, []int{k, v})
	}
	sort.Slice(freqints, func(i, j int) bool {
		return freqints[i][1] > freqints[j][1]
	})
	out := make([]int, k)
	for i := 0; i < k; i++ {
		out[i] = freqints[i][0]
	}
	return out
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
