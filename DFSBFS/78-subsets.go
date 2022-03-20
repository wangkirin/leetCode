package main

func main() {
	subsets([]int{1, 2, 3})
}

//BFS
func subsets(nums []int) [][]int {
	ans := make([][]int, 1)
	ans[0] = []int{}
	for _, x := range nums {
		for _, arr := range ans {
			a := make([]int, len(arr))
			copy(a, arr)
			ans = append(ans, append(a, x))
		}
	}
	return ans
}

//func subsets(nums []int) [][]int {
//	ans := [][]int{}
//	var dfs = func(nums []int, track []int) {}
//	track := []int{}
//	dfs = func(nums []int, track []int) {
//		if len(nums) == len(track) {
//			ans = append(ans, append([]int{}, track...))
//			return
//		}
//		for i := 0; i < len(nums); i++ {
//
//		}
//	}
//	dfs(nums, track)
//	return ans
//
//}

//回溯算法
//func subsets(nums []int) (ans [][]int) {
//	set := []int{}
//	var dfs func(int)
//	dfs = func(cur int) {
//		if cur == len(nums) {
//			ans = append(ans, append([]int(nil), set...))
//			return
//		}
//		set = append(set, nums[cur])
//		dfs(cur + 1)
//		set = set[:len(set)-1]
//		dfs(cur + 1)
//	}
//	dfs(0)
//	return
//}
