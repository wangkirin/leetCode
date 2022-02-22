package main

import "sort"

func main() {

}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	out := [][]int{intervals[0]}
	for i := 0; i < len(intervals); i++ {
		if out[len(out)-1][1] >= intervals[i][0] {
			out[len(out)-1][1] = max(out[len(out)-1][1], intervals[i][1])
		} else {
			out = append(out, intervals[i])
		}
	}
	return out

}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
