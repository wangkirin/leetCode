package main

func main() {

}

//func findAnagrams(s string, p string) []int {
//	ans := []int{}
//	m := make(map[byte]int)
//	window := make(map[byte]int)
//	for i := 0; i < len(p); i++ {
//		m[p[i]]++
//	}
//	l, valid := 0, 0
//	for r := 0; r < len(s); r++ {
//		if _, ok := m[s[r]]; ok {
//			window[s[r]]++
//			if window[s[r]] == m[s[r]] {
//				valid++
//			}
//		}
//		for r-l >= len(p) {
//			if valid == len(p) {
//				ans = append(ans, l)
//			}
//			l++
//			if _, ok := m[s[l]]; ok {
//				if window[s[l]] == m[s[l]] {
//					valid--
//				}
//				window[s[l]]--
//			}
//		}
//	}
//	return ans
//}
