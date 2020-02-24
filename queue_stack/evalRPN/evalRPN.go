package evalRPN

import (
	"container/list"
	"strconv"
)

//根据逆波兰表示法，求表达式的值。
//
//有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
//
//说明：
//
//整数除法只保留整数部分。
//给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
//示例 1：
//
//输入: ["2", "1", "+", "3", "*"]
//输出: 9
//解释: ((2 + 1) * 3) = 9
//示例 2：
//
//输入: ["4", "13", "5", "/", "+"]
//输出: 6
//解释: (4 + (13 / 5)) = 6
//示例 3：
//
//输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
//输出: 22
//解释:
//((10 * (6 / ((9 + 3) * -11))) + 17) + 5
//= ((10 * (6 / (12 * -11))) + 17) + 5
//= ((10 * (6 / -132)) + 17) + 5
//= ((10 * 0) + 17) + 5
//= (0 + 17) + 5
//= 17 + 5
//= 22

//No.150

//从左往右依次读取表达式，如果是数字则将该数字压栈，如果是符号，则将之前的两个数字出栈，做计算后，将计算结果压栈，直到表达式读取结束。栈中剩下的一个数就是计算结果。
func evalRPN(tokens []string) int {
	result:=0
	stack:=list.New()
	for i:=0;i<len(tokens) ;  i++{
		num,err:=strconv.Atoi(tokens[i])
		if err==nil {
			stack.PushBack(num)
		}else{
			switch tokens[i] {
			case "+":
				cal1:=stack.Back().Value.(int)
				stack.Remove(stack.Back())
				cal2:=stack.Back().Value.(int)
				stack.Remove(stack.Back())
				cal:=cal1+cal2
				stack.PushBack(cal)
			case "-":
				cal1:=stack.Back().Value.(int)
				stack.Remove(stack.Back())
				cal2:=stack.Back().Value.(int)
				stack.Remove(stack.Back())
				cal:=cal2-cal1
				stack.PushBack(cal)
			case "*":
				cal1:=stack.Back().Value.(int)
				stack.Remove(stack.Back())
				cal2:=stack.Back().Value.(int)
				stack.Remove(stack.Back())
				cal:=cal1*cal2
				stack.PushBack(cal)
			case "/":
				cal1:=stack.Back().Value.(int)
				stack.Remove(stack.Back())
				cal2:=stack.Back().Value.(int)
				stack.Remove(stack.Back())
				cal:=cal2/cal1
				stack.PushBack(cal)
			}
		}
	}
	if stack.Len()>0 {
		result=stack.Back().Value.(int)
	}
	return result
}
