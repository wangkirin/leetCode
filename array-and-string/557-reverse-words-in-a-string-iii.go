package main

func main() {

}

func reverseWordsII(s string) string {
	runes := []rune(s)
	start := 0
	for i := 0; i < len(runes); i++ {
		check := false
		if runes[i] == ' ' {
			check = true
		}
		if !check && i == len(runes)-1 {
			check = true
			i = i + 1
		}
		if check {
			if i > start {
				for j := 0; j <= (i-1-start)/2; j++ {
					runes[i-1-j], runes[start+j] = runes[start+j], runes[i-1-j]
				}
			}
			start = i + 1
		}
	}
	return string(runes)
}
