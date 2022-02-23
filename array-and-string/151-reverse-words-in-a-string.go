package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseWords("  hello world  "))
	fmt.Println(reverseWords("a good   example"))
}

func reverseWords(s string) string {
	//1.使用双指针删除冗余的空格
	slowIndex, fastIndex := 0, 0
	b := []byte(s)
	//删除头部冗余空格
	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' ' {
		fastIndex++
	}
	//删除单词间冗余空格
	for ; fastIndex < len(b); fastIndex++ {
		if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}
	//删除尾部冗余空格
	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' {
		b = b[:slowIndex-1]
	} else {
		b = b[:slowIndex]
	}
	//2.反转整个字符串
	reverse(&b, 0, len(b)-1)
	//3.反转单个单词  i单词开始位置，j单词结束位置
	i := 0
	for i < len(b) {
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {
		}
		reverse(&b, i, j-1)
		i = j
		i++
	}
	return string(b)
}

func reverse(b *[]byte, left, right int) {
	for left < right {
		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}
}

//func reverseWords(s string) string {
//	// 双指针翻转字符串
//	sbytes := []byte(s)
//	r := revers(sbytes)
//	//处理前空格
//	for i := 0; i < len(r); i++ {
//		if byte(' ') != r[i] {
//			r = r[i:]
//			break
//		}
//	}
//	//处理后空格
//	for i := len(r) - 1; i >= 0; i-- {
//		if r[i] != byte(' ') {
//			r = r[0 : i+1]
//		}
//	}
//	for k, v := range r {
//
//	}
//
//	return
//}
//func revers(in []byte) []byte {
//	left := 0
//	right := len(in) - 1
//	for left < right {
//		temp := in[left]
//		in[left] = in[right]
//		in[right] = temp
//		left++
//		right--
//	}
//	return in
//}

//1.普通解法
//func reverseWords(s string) string {
//	out := ""
//	s = strings.TrimSpace(s)
//	frag := strings.Fields(s)
//	for i := len(frag) - 1; i >= 0; i-- {
//		out = out + frag[i]
//		if i != 0 {
//			out += " "
//		}
//	}
//	return out
//}
