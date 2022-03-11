package main

import "strings"

func main() {

}
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	strings.ReplaceAll(s, " ", "")
	left := 0
	r := []byte(s)
	right := len(r) - 1
	for left < right {
		if !isalnum(r[left]) {
			left++
			continue
		}
		if !isalnum(r[right]) {
			right--
			continue
		}
		if r[left] != r[right] {
			return false
		}
		left++
		right--

	}
	return true
}

//判断是否字母或者数字
func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
