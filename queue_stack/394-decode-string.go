package main

func main() {

}
func decodeString(s string) string {
	out := ""
	ss := []rune(s)
	stack := []rune{}
	for len(ss) > 0 {
		if ss[0] != ']' {
			stack = append(stack, ss[0])
			ss = ss[1:]
		} else {

		}

	}
	return out
}
