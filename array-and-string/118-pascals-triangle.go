package main

func main() {
	println(generate(5))
}

func generate(numRows int) [][]int {
	ans := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		lastRow, nextRow := ans[len(ans)-1], make([]int, i+1)
		nextRow[0], nextRow[i] = 1, 1
		for j := 1; j < i; j++ {
			nextRow[j] = lastRow[j-1] + lastRow[j]
		}
		ans = append(ans, nextRow)
	}
	return ans
}

//func generate(numRows int) [][]int {
//	out := [][]int{}
//	for i := 1; i <= numRows; i++ {
//		out = append(out, line(i))
//	}
//	return out
//}
//func line(lineNum int) []int {
//	out := make([]int, lineNum)
//	if lineNum == 1 {
//		out[0] = 1
//		return out
//	}
//	if lineNum == 2 {
//		out[0] = 1
//		out[1] = 1
//		return out
//	}
//	out[0] = 1
//	out[lineNum-1] = 1
//	last := line(lineNum - 1)
//	for i := 1; i < lineNum-2; i++ {
//		fmt.Println("i=", i)
//
//		out[i] = last[i-1] + last[i]
//		fmt.Println("out i", out[i])
//	}
//	return out
//}
