package main

func main() {

}

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	maxa := min(height[l], height[r]) * (r - l)
	for l < r {
		if min(height[l], height[r])*(r-l) > maxa {
			maxa = min(height[l], height[r]) * (r - l)
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxa
}
