package main

func main() {

}

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		index := start + (end-start)/2
		if nums[index] == target {
			return index
		}
		if nums[index] > target {
			end = index - 1
		} else {
			start = index + 1
		}
	}
	return -1
}
