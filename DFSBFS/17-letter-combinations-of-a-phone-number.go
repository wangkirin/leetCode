package main

func main() {

}
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	letters := [][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
		{'j', 'k', 'l'},
		{'m', 'n', 'o'},
		{'p', 'q', 'r', 's'},
		{'t', 'u', 'v'},
		{'w', 'x', 'y', 'z'},
	}
	digitsbytes := []byte(digits)
	ans := []string{}
	for _, digit := range digitsbytes {
		if len(ans) == 0 {
			for _, v := range letters[digit-'2'] {
				ans = append(ans, string(v))
			}
		} else {
			lenque := len(ans)
			for i := 0; i < lenque; i++ {
				for _, letter := range letters[digit-'2'] {
					anbyte := []byte(ans[i])
					anbyte = append(anbyte, letter)
					ans = append(ans, string(anbyte))
				}
			}
			ans = ans[lenque:]
		}
	}
	return ans
}
