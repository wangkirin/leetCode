package main

func main() {

}

// 前缀和
func canEat(candiesCount []int, queries [][]int) []bool {
	preSum := make([]int, len(candiesCount)+1)
	preSum[0] = 0
	for i := 0; i < len(candiesCount); i++ {
		preSum[i+1] = preSum[i] + candiesCount[i]
	}
	ans := make([]bool, len(queries))
	for i := 0; i < len(queries); i++ {
		query := queries[i]
		t, day, limit := query[0], query[1], query[2]
		if preSum[t] < (day+1)*limit && preSum[t+1] >= (day+1) {
			ans[i] = true
		} else {
			ans[i] = false
		}
	}
	return ans
}
