package main

func main() {

}

func mySqrt(x int) int {
	start := 0
	end := x
	for start <= end {
		mid := start + (end-start)/2
		squre := mid * mid
		if squre == x {
			return mid
		}
		if squre > x {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return end
}
