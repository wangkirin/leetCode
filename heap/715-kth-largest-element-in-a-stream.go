package main

import (
	"container/heap"
	"sort"
)

func main() {

}

//优先队列
type KthLargest struct {
	sort.IntSlice
	k int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, val := range nums {
		kl.Add(val)
	}
	return kl
}

func (kl *KthLargest) Push(v interface{}) {
	kl.IntSlice = append(kl.IntSlice, v.(int))
}

func (kl *KthLargest) Pop() interface{} {
	a := kl.IntSlice
	v := a[len(a)-1]
	kl.IntSlice = a[:len(a)-1]
	return v
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(kl, val)
	if kl.Len() > kl.k {
		heap.Pop(kl)
	}
	return kl.IntSlice[0]
}

////堆
//type KthLargest struct {
//	heap *[]int
//	k    int
//}
//
//func Constructor(k int, nums []int) KthLargest {
//	h := KthLargest{}
//	h.heap = &[]int{}
//	heap.Init(&h)
//	h.k = k
//	for _, num := range nums {
//		heap.Push(&h, num)
//	}
//	return h
//}
//
//func (this *KthLargest) Add(val int) int {
//	this.Push(val)
//	for i := 0; i < this.k-1; i++ {
//		heap.Pop(this)
//	}
//	fmt.Println(this.heap)
//	return this.Peek()
//}
//
////Less  小于就是小跟堆，大于号就是大根堆
//func (h *KthLargest) Less(i, j int) bool { return (*h.heap)[i] > (*h.heap)[j] }
//func (h *KthLargest) Swap(i, j int)      { (*h.heap)[i], (*h.heap)[j] = (*h.heap)[j], (*h.heap)[i] }
//func (h *KthLargest) Len() int           { return len(*h.heap) }
//func (h *KthLargest) Push(x interface{}) {
//	*h.heap = append(*h.heap, x.(int))
//}
//func (h *KthLargest) Pop() interface{} {
//	t := len(*h.heap) - 1
//	*h.heap = (*h.heap)[:len(*h.heap)-1]
//	return t
//}
//func (h *KthLargest) Peek() int {
//	return (*h.heap)[0]
//}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
