package main

func main() {

}

func maxArea(height []int) int {
	maxa := 0
	l, r := 0, len(height)-1
	for l < r {
		maxa = max(min(height[r], height[l])*(r-l), maxa)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}

	}
	return maxa
}
