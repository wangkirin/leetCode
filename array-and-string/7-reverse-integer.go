package main

import "math"

func main() {
	reverseq(123)
}

//
func reverse(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return
}

func reverseq(x int) int {

	smallzero := false
	if x < 0 {
		smallzero = true
		x *= -1
	}
	out := 0
	for x > 10 {
		out = out*10 + (x % 10)
		x = x / 10

	}
	out = out*10 + x
	if smallzero {
		return out * -1
	}
	if out < math.MinInt32/10 || out > math.MaxInt32/10 {
		return 0
	}
	return out

}
