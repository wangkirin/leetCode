package main

import "sort"

func main() {

}
func isAnagram(s string, t string) bool {
	bytess := []byte(s)
	bytest := []byte(t)
	if len(s) != len(t) {
		return false
	}
	sort.Slice(bytest, func(i, j int) bool {
		return bytest[i] < bytest[j]
	})
	sort.Slice(bytess, func(i, j int) bool {
		return bytess[i] < bytess[j]
	})
	for i := 0; i < len(bytest); i++ {
		if bytess[i] != bytest[i] {
			return false
		}
	}
	return true
}
