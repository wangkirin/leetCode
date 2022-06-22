package main

import "fmt"

func main() {

}

func maxArea(height []int) int {
	maxa := 0
	l, r := 0, len(height)-1
	for l < r {
		fmt.Printf("l=%d,r=%d,height[l]=%d,height[r]=%d,maxa=%d \n", l, r, height[l], height[r], maxa)
		maxa = max(min(height[r], height[l])*(r-l), maxa)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}

	}
	return maxa
}

func maxArea1(height []int) int {
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
