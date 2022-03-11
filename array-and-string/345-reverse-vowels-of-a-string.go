package main

import "strings"

func main() {

}

//双指针
func reverseVowels(s string) string {
	strings.ReplaceAll(s, " ", "")
	left := 0
	r := []byte(s)
	right := len(r) - 1
	for left < right {
		if !isVowels(r[left]) {
			left++
			continue
		}
		if !isVowels(r[right]) {
			right--
			continue
		}
		r[left], r[right] = r[right], r[left]
		left++
		right--

	}
	return string(r)
}

//判断是否字母或者数字
func isVowels(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' || ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U'
}
