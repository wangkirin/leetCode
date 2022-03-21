package main

func main() {

}
func generateParenthesis(n int) []string {

}

//DFS
//func generateParenthesis(n int) []string {
//	res := []string{}
//
//	var dfs func(lRemain int, rRemain int, path string)
//	dfs = func(lRemain int, rRemain int, path string) {
//		if 2*n == len(path) {
//			res = append(res, path)
//			return
//		}
//		if lRemain > 0 {
//			dfs(lRemain-1, rRemain, path+"(")
//		}
//		if lRemain < rRemain {
//			dfs(lRemain, rRemain-1, path+")")
//		}
//	}
//
//	dfs(n, n, "")
//	return res
//}
//
//作者：xiao_ben_zhu
//链接：https://leetcode-cn.com/problems/generate-parentheses/solution/shou-hua-tu-jie-gua-hao-sheng-cheng-hui-su-suan-fa/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请
//注明出处。
