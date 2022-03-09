package main

import "fmt"

func main() {
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'd'))
}

func nextGreatestLetter(letters []byte, target byte) byte {
	l := 0
	r := len(letters) - 1
	for l < r {
		mid := l + (r-l)/2
		if letters[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return letters[l%len(letters)]
}
