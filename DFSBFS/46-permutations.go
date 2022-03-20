package main

func main() {

}
func permute(nums []int) [][]int {
	ans := [][]int{}
	track := []int{}
	visited := make(map[int]bool)
	var dfs func()
	dfs = func() {
		if len(track) == len(nums) {
			ans = append(ans, append([]int{}, track...))
			//原来使用这行，有错误
			//因为该 track 变量存的是地址引用，结束当前递归时，将它加入 res 后，该算法还要进入别的递归分支继续搜索，还要继续将这个 track 传给别的递归调用，它所指向的内存空间还要继续被操作，所以 res 中的 path 的内容会被改变，这就不对。
			//所以要弄一份当前的拷贝，放入 res，这样后续对 track 的操作，就不会影响已经放入 res 的内容。
			//ans = append(ans, track)
			return
		}
		for i := 0; i < len(nums); i++ {
			if visited[nums[i]] {
				continue
			}
			visited[nums[i]] = true
			track = append(track, nums[i])

			dfs()
			visited[nums[i]] = false
			track = track[:len(track)-1]
		}
	}
	dfs()
	return ans
}
