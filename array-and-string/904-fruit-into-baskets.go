package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(totalFruit([]int{0, 1, 1, 4, 3}))
}

//优化
func totalFruit(fruits []int) int {
	l, maxL := 0, 0
	window := make(map[int]int)
	for r := 0; r < len(fruits); r++ {
		window[fruits[r]]++
		for len(fruits) > 2 && l < r {
			if window[fruits[l]] == 1 {
				delete(window, fruits[l])
			}
			window[fruits[l]]--
			l++
		}
		maxL = max(maxL, r-l+1)
	}
	return maxL
}

func totalFruit1(fruits []int) int {
	l, maxL := 0, 0
	window := make(map[int]int)
	for r := 0; r < len(fruits); r++ {
		_, ok := window[fruits[r]]
		if len(window) == 2 && !ok {
			tmp := [][2]int{}
			for k, v := range window {
				tmp = append(tmp, [2]int{k, v})
			}
			sort.Slice(tmp, func(i, j int) bool {
				return tmp[i][1] < tmp[j][1]
			})

			delete(window, tmp[0][0])
			l = tmp[0][1] + 1
		}
		window[fruits[r]] = r
		maxL = max(maxL, r-l+1)
	}
	return maxL
}
func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
