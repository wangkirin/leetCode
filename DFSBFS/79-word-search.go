package main

func main() {

}
func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	wordBytes := []byte(word)
	dfs := func(r, c int, words []byte) bool {
		if r < 0 || c < 0 || r >= m || c >= n || board[r][c] != words[0] {
			return false
		}
		if dfs(r-1, c, words[0:]) || dfs(r+1, c, words[0:]) || dfs(r, c-1, words[0:]) || dfs(r, c+1, words[0:]) {
			return true
		}
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == wordBytes[0] {
				if dfs(i, j, wordBytes[0:]) {
					return true
				}

			}
		}
	}
	return false
}
