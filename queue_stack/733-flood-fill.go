package main

func main() {

}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	dfs(image, sr, sc, image[sr][sc], newColor)
	return image
}

func dfs(grid [][]int, r, c int, oldcolor int, newColor int) {
	h, w := len(grid), len(grid[0])
	if r < 0 || r >= h || c < 0 || c >= w || grid[r][c] != oldcolor || newColor == oldcolor {
		return
	}
	grid[r][c] = newColor
	dfs(grid, r-1, c, oldcolor, newColor)
	dfs(grid, r+1, c, oldcolor, newColor)
	dfs(grid, r, c-1, oldcolor, newColor)
	dfs(grid, r, c+1, oldcolor, newColor)
}
