package main

import "fmt"

func main() {

}

func intervalIntersection(A [][]int, B [][]int) [][]int {
	res := [][]int{}
	i, j := 0, 0
	for i < len(A) && j < len(B) {
		start := max(A[i][0], B[j][0])
		end := min(A[i][1], B[j][1])
		if start <= end {
			res = append(res, []int{start, end})
		}
		if A[i][1] < B[j][1] {
			i++
		} else {
			j++
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//合到一个数组里，解决不了区间问题
func intervalIntersection1(firstList [][]int, secondList [][]int) [][]int {
	max1 := firstList[len(firstList)-1][1]
	max2 := secondList[len(secondList)-1][1]
	max := max(max1, max2)
	arr := make([]int, max+1)
	for _, ints := range firstList {
		for i := ints[0]; i <= ints[1]; i++ {
			arr[i]++
		}
	}
	for _, ints := range secondList {
		for i := ints[0]; i <= ints[1]; i++ {
			arr[i]++
		}
	}
	fmt.Println(arr)
	ans := [][]int{}
	start := false
	starti, endi := 0, 0
	for k, v := range arr {
		fmt.Printf("k=%d,v=%d,starti=%d, endi=%d \n", k, v, starti, endi)
		if v == 2 {
			if !start {
				start = true
				starti, endi = k, k
			} else {
				endi = k
			}
			continue
		}
		if start {
			ans = append(ans, []int{starti, endi})
			start = false
		}
	}
	return ans
}
