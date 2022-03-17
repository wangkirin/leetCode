package main

func main() {

}
func permute(nums []int) [][]int {
	ans := [][]int{}
	round := 1

	for round <= len(nums) {
		if len(ans) == 0 {
			for _, v := range nums {
				ans = append(ans, []int{v})
			}
		}
		anslen := len(ans)

		round++
	}
	return ans
}

func dfs(nums []int) []int {

}
