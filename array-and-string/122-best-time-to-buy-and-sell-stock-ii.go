package main

func main() {

}

func maxProfit(prices []int) int {
	sum := 0
	start, end := 0, 0
	if len(prices) == 1 {
		return 0
	}
	if len(prices) == 2 {
		if prices[1] > prices[0] {
			return prices[1] - prices[0]
		} else {
			return 0
		}
	}
	for r := 1; r < len(prices)-1; r++ {
		if prices[r] < prices[r-1] && prices[r] < prices[r+1] {
			start = prices[r]
		}
		if prices[r] > prices[r-1] && prices[r] > prices[r+1] {
			end = r
			sum += (prices[end] - prices[start])
		}
	}
	if start == end && prices[len(prices)-1] > prices[0] {
		sum = prices[len(prices)-1] - prices[0]
	}
	return sum
}
