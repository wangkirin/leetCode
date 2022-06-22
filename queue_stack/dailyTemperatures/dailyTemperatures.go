package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}))
}

//单调栈解法

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := []int{}
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			ans[i] = 0

		} else {
			ans[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return ans
}

func dailyTemperatures1(temperatures []int) []int {
	out := make([]int, len(temperatures))
	stack := []int{}

	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			out[i] = 0
		} else {
			out[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return out
}

//No.739
//根据每日 气温 列表，请重新生成一个列表，对应位置的输入是你需要再等待多久温度才会升高超过该日的天数。如果之后都不会升高，请在该位置用 0 来代替。
//
//例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。
//
//提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数

//func dailyTemperatures(T []int) []int {
//	res := make([]int, len(T), len(T))
//	stack := list.New()
//	for i, num := range T {
//		for stack.Len() != 0 && T[stack.Front().Value.(int)] < num {
//			ind := stack.Remove(stack.Front()).(int)
//			res[ind] = i-ind
//		}
//		stack.PushFront(i)
//	}
//	return res
//}

//https://leetcode-cn.com/problems/daily-temperatures/solution/golangshi-yong-containerlistbao-by-yang7y7y7/

//func dailyTemperatures(T []int) []int {
//	Out:=make([]int,len(T))
//	stack:=[]int{}
//	for i:= len(T)-1;i>=0 ;i--{
//		lt:=len(stack)
//		if lt==0 {
//			Out[i]=0
//			stack=append(stack,i)
//			continue
//		}
//		if T[stack[lt-1]]>T[i] {
//			Out[i]= stack[lt-1]-i
//			stack=append(stack,i)
//
//		}else {
//			for j:=lt-1;j>=0 ;j--  {
//				if (T[stack[j]]>T[i]){
//					Out[i]= stack[j]-i
//					break
//				}else {
//					stack=stack[0:len(stack)-1]
//					stack=append(stack,i)
//				}
//			}
//
//		}
//	}
//	return Out
//}

//// 老土解法：暴力解
//func dailyTemperatures(T []int) []int {
//	Out:=make([]int,len(T))
//	for i:=0;i<len(T); i++{
//		for j:=i+1;j<len(T);j++  {
//			if T[j]>T[i] {
//				Out[i]=(j-i)
//				break
//			}
//		}
//	}
//	return Out
//
//}
