package main

func main() {

}

//求所有单调递增区间
func maxProfit(prices []int) int {
	sum := 0
	if len(prices) == 1 {
		return 0
	}
	for r := 1; r < len(prices); r++ {
		if prices[r] > prices[r-1] {
			sum += (prices[r] - prices[r-1])
		}
	}
	return sum
}
