package main

func main() {

}

func gameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])
	bak := make([][]int, m)
	for i := 0; i < m; i++ {
		bak[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			bak[i][j] = board[i][j]
		}
	}
	searchlist := [8][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			count := 0
			for _, s := range searchlist {
				r, c := s[0], s[1]
				if i+r < 0 || i+r > m-1 || j+c < 0 || j+c > n-1 {
					continue
				}
				count += bak[i+r][j+c]
			}
			if bak[i][j] == 0 {
				if count == 3 {
					board[i][j] = 1
				} else {
					board[i][j] = 0
				}
			} else if bak[i][j] == 1 {
				if count == 3 || count == 2 {
					board[i][j] = 1
				} else {
					board[i][j] = 0
				}
			}
		}
	}
}
