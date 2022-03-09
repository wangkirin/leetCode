package main

import "fmt"

func main() {
	isPerfectSquare(16)
}
func isPerfectSquare(num int) bool {
	left := 0
	right := num
	for left <= right {
		mid := left + (right-left)/2
		fmt.Println(mid)
		if mid*mid == num {
			return true
		}
		if mid*mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
