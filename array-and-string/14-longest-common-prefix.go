package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
}
func longestCommonPrefix(strs []string) string {
	out := ""
	for i := 0; i < len(strs[0]); i++ {
		first := []byte(strs[0])
		word := first[i]
		breakFlag := false
		for _, v := range strs {
			if len(v) < i+1 || (len(v) > i && v[i] != word) {
				breakFlag = true
				break
			}
		}
		if breakFlag {
			break
		}
		out = out + string([]byte{word})
	}
	return out
}
