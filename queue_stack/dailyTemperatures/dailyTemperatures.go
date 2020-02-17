package dailyTemperatures

import "fmt"

//No.739
//根据每日 气温 列表，请重新生成一个列表，对应位置的输入是你需要再等待多久温度才会升高超过该日的天数。如果之后都不会升高，请在该位置用 0 来代替。
//
//例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。
//
//提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数


func dailyTemperatures(T []int) []int {
	Out:=make([]int,len(T))
	for i:=0;i<len(T); i++{
		for j:=i+1;j<len(T);j++  {
			fmt.Printf("i=%d, j=%d,len(T)=%d\n",i,j,len(T))
			if T[j]>T[i] {
				Out[i]=(j-i)
				break
			}
		}
	}
	return Out

}
