package main

import "sort"

func main() {

}

type Node struct {
	father    int
	available int
}

type BikeManager struct {
	nodes map[int]*Node
	cap   int
}

func Constructor(preNode []int, capacicty int) BikeManager {
	b := BikeManager{cap: capacicty}
	b.nodes[0] = &Node{
		father:    -1,
		available: 1000,
	}
	for i := 1; i < len(preNode); i++ {
		b.nodes[i] = &Node{
			father:    preNode[i],
			available: b.cap / 2,
		}
	}
	return b
}

func (b *BikeManager) rentBikes(nodeid int, num int) int {
	if b.nodes[nodeid].available >= num {
		b.nodes[nodeid].available = b.nodes[nodeid].available - num
		return b.nodes[nodeid].available
	} else {
		left := num - b.nodes[nodeid].available
		k := b.nodes[nodeid].father
		for {
			if k == 0 {
				break
			}
			if b.nodes[k].available >= left {
				b.nodes[k].available = b.nodes[k].available - left
				break
			} else {
				left = left - b.nodes[k].available
				b.nodes[k].available = 0
				k = b.nodes[k].father

			}
		}
	}
	return 0
}

func (b *BikeManager) returnBikes(nodeid int, num int) int {
	if b.nodes[nodeid].available+num <= b.cap {
		b.nodes[nodeid].available = b.nodes[nodeid].available + num
		return b.nodes[nodeid].available
	} else {
		more := b.nodes[nodeid].available + num - b.cap
		k := b.nodes[nodeid].father
		for {
			if k == 0 {
				break
			}
			if b.nodes[k].available+more < b.cap {
				b.nodes[k].available = b.nodes[k].available + more
				break
			} else {
				more = more - b.nodes[k].available
				b.nodes[k].available = b.cap
				k = b.nodes[k].father
			}
		}
	}
	return b.cap
}

func (b *BikeManager) reset() int {
	count := 0
	for i := 1; i < len(b.nodes); i++ {
		if b.nodes[i].available == 0 || b.nodes[i].available == b.cap {
			b.nodes[i].available = b.cap / 2
			count++
		}
	}
	return count
}

func (b *BikeManager) getTop5Nodes() []int {
	ans := []int{}
	keyAvalibleSlices := [][2]int{}
	for k, node := range b.nodes {
		keyAvalibleSlices = append(keyAvalibleSlices, [2]int{k, node.available})
	}

	sort.Slice(keyAvalibleSlices, func(i, j int) bool {
		if keyAvalibleSlices[i][1] == keyAvalibleSlices[j][1] {
			return keyAvalibleSlices[i][0] < keyAvalibleSlices[j][0]
		}
		return keyAvalibleSlices[i][1] > keyAvalibleSlices[j][1]
	})

	if len(keyAvalibleSlices) < 5 {
		for i := 0; i < len(keyAvalibleSlices); i++ {
			ans = append(ans, keyAvalibleSlices[i][0])
		}
	} else {
		for i := 0; i < 5; i++ {
			ans = append(ans, keyAvalibleSlices[i][0])
		}
	}
	return ans
}
