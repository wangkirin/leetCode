package main

import (
	"container/heap"
	"sort"
)

func main() {

}

func topKFrequent(words []string, k int) []string {
	wordsMap := make(map[string]int)
	for _, word := range words {
		count := wordsMap[word]
		wordsMap[word] = count + 1
	}
	pq := make(PriorityQueue, len(wordsMap))
	i := 0
	for k, v := range wordsMap {
		pq[i] = &Item{
			value:    k,
			priority: v,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)
	ans := []string{}
	for i := 0; i < k; i++ {
		item := heap.Pop(&pq).(*Item)
		ans = append(ans, item.value)
	}
	return ans
}

type wordCounts struct {
	value string
	count int
}

// Item 是优先队列中包含的元素。
type Item struct {
	value    string // 元素的值，可以是任意字符串。
	priority int    // 元素在队列中的优先级。
	// 元素的索引可以用于更新操作，它由 heap.Interface 定义的方法维护。
	index int // 元素在堆中的索引。
}

// 一个实现了 heap.Interface 接口的优先队列，队列中包含任意多个 Item 结构。
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// 我们希望 Pop 返回的是最大值而不是最小值，
	// 因此这里使用大于号进行对比。
	if pq[i].priority == pq[j].priority {
		return pq[i].value < pq[j].value
	}
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // 为了安全性考虑而做的设置
	*pq = old[0 : n-1]
	return item
}

// 更新函数会修改队列中指定元素的优先级以及值。
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// 普通解法
func topKFrequent1(words []string, k int) []string {
	wordsMap := make(map[string]int)
	for _, word := range words {
		count := wordsMap[word]
		wordsMap[word] = count + 1
	}
	wcs := []wordCounts{}
	for k, v := range wordsMap {
		wcs = append(wcs, wordCounts{
			value: k,
			count: v,
		})
	}
	sort.Slice(wcs, func(i, j int) bool {
		if wcs[i].count == wcs[j].count {
			return wcs[i].value < wcs[j].value
		} else {
			return wcs[i].count > wcs[j].count
		}
	})
	ans := []string{}
	for i := 0; i < k; i++ {
		ans = append(ans, wcs[i].value)
	}
	return ans
}
