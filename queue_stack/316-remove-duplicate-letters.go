package main

import "fmt"

func main() {

}

func removeDuplicateLetters(s string) string {
	set := make(map[byte]int)
	inStack := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		set[s[i]]++
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		set[s[i]]--
		if inStack[s[i]] {
			continue
		}
		for len(stack) > 0 &&  && stack[len(stack)-1] > s[i] {
			if set[stack[len(stack)-1]] == 0 {
				break
			}
			inStack[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, s[i])
		inStack[s[i]] = true
	}
	ans := ""
	for _, b := range stack {
		ans = fmt.Sprintf("%s%c", ans, b)
	}
	return ans
}
