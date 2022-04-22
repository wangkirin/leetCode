package main

func main() {

}
func calculate(s string) int {
	sbytes := []rune(s)
	stack := []int{}
	op := '+'
	num := 0
	for i := 0; i < len(sbytes); i++ {
		if sbytes[i] >= '0' && sbytes[i] <= '9' {
			count := int(sbytes[i] - '0')
			num = num*10 + count
		}
		if (!(sbytes[i] >= '0' && sbytes[i] <= '9') && sbytes[i] != ' ') || (i == len(sbytes)-1) {
			switch op {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -1*num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			op = sbytes[i]
			num = 0
		}
	}
	ans := 0
	for _, s := range stack {
		ans += s
	}
	return ans
}
