package main

import "fmt"

func main() {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
}
func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	wordBytes := []byte(word)
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var dfs func(board [][]byte, r, c, i int) bool
	dfs = func(board [][]byte, r, c, i int) bool {
		if i == len(word) {
			return true
		}

		if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) || visited[r][c] || board[r][c] != wordBytes[i] {
			return false
		}
		visited[r][c] = true
		if dfs(board, r-1, c, i+1) || dfs(board, r+1, c, i+1) || dfs(board, r, c-1, i+1) || dfs(board, r, c+1, i+1) {
			return true
		} else {
			visited[r][c] = false
			return false
		}

	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == wordBytes[0] {
				if dfs(board, i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}
