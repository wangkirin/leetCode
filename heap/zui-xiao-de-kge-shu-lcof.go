package main

import (
	"container/heap"
	"fmt"
)

func main() {
	//arr := []int{3, 2, 1}
	//fmt.Println(getLeastNumbers(arr, 2))
	fmt.Println(getLeastNumbers([]int{0, 1, 2, 1}, 1))
}

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}
	//方法三  建堆，大根堆
	d := &heapInt{}
	for _, v := range arr {
		if d.Len() < k {
			heap.Push(d, v)
		} else {
			if d.Peek() > v {
				heap.Pop(d)
				heap.Push(d, v)
			}
		}
	}
	return *d
}

type heapInt []int

//Less  小于就是小跟堆，大于号就是大根堆
func (h *heapInt) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *heapInt) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *heapInt) Len() int           { return len(*h) }
func (h *heapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *heapInt) Pop() interface{} {
	t := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return t
}
func (h *heapInt) Peek() int {
	return (*h)[0]
}

//func getLeastNumbers(arr []int, k int) []int {
//	myheap := IntHeap{}
//	heap.Init(&myheap)
//	for _, v := range arr {
//		myheap.Push(v)
//	}
//	out := []int{}
//
//	for i := 0; i < k; i++ {
//		fmt.Println("i=", i)
//		fmt.Println("myHeap=", myheap)
//		out = append(out, heap.Pop(&myheap).(int))
//	}
//	return out
//}
//
//type IntHeap []int
//
//func (h IntHeap) Len() int           { return len(h) }
//func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
//func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//func (h *IntHeap) Push(x interface{}) {
//	// Push 和 Pop 使用 pointer receiver 作为参数，
//	// 因为它们不仅会对切片的内容进行调整，还会修改切片的长度。
//	*h = append(*h, x.(int))
//}
//
//func (h *IntHeap) Pop() interface{} {
//	old := *h
//	n := len(old)
//	x := old[n-1]
//	*h = old[0 : n-1]
//	return x
//}
