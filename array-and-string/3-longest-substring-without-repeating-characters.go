package main

import "fmt"

func main() {
	//lengthOfLongestSubstring("abcabcbb")
	lengthOfLongestSubstring("bbb")
}
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	l := 0
	bs := []byte(s)
	hash := make(map[byte]int)
	longest := 1

	for r := 0; r < len(bs); r++ {
		fmt.Println("r=", r, bs[r])
		if _, ok := hash[bs[r]]; !ok {
			hash[bs[r]] = 1
			fmt.Println("longest=", longest)
		} else {
			if (r - l) > longest {
				longest = r - l
			}
			hash = make(map[byte]int)
			hash[bs[r]] = 1
			l++
		}
	}
	return longest
}

//双指针解法
func lengthOfLongestSubstring(s string) int {
	res := 0
	left, right := 0, 0
	for right < len(s) {
		for i := left; i < right; i++ {
			if s[i] == s[right] {
				res = max(res, right-left)
				left = i + 1
				break
			}
		}
		right++
	}
	res = max(res, right-left)
	return res
}
