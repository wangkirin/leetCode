package main

import "fmt"

func main() {
	//lengthOfLongestSubstring("abcabcbb")
	//lengthOfLongestSubstring("bbb")
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
}

//自己做优化
func lengthOfLongestSubstring(s string) int {
	l, count := 0, 0
	windowCount := make(map[byte]int)
	for r := 0; r < len(s); r++ {
		if _, ok := windowCount[s[r]]; !ok {
			windowCount[s[r]] = r
			count = max(count, r-l+1)
		} else {
			delete(windowCount, s[l])
			l = windowCount[s[r]] + 1
			windowCount[s[r]] = 1
		}
	}
	return count
}

//自己做1
func lengthOfLongestSubstring3(s string) int {
	l, count := 0, 0
	windowCount := make(map[byte]int)
	for r := 0; r < len(s); r++ {
		if _, ok := windowCount[s[r]]; !ok {
			windowCount[s[r]] = 1
			count = max(count, r-l+1)
		} else {
			for s[l] != s[r] && l < r {
				delete(windowCount, s[l])
				l++
			}
			l++
			windowCount[s[r]] = 1
		}
	}
	return count
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func lengthOfLongestSubstring2(s string) int {
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
func lengthOfLongestSubstring1(s string) int {
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
