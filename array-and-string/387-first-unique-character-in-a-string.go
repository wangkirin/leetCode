package main

func main() {

}

func firstUniqChar(s string) int {
	hash := make(map[rune]int)
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		hash[rs[i]]++
	}
	for j := 0; j < len(rs); j++ {
		if hash[rs[j]] == 1 {
			return j
		}
	}
	return -1
}

//可以对方法一进行修改，使得第二次遍历的对象从字符串变为哈希映射。
//具体地，对于哈希映射中的每一个键值对，键表示一个字符，值表示它的首次出现的索引（如果该字符只出现一次）或者 -1（如果该字符出现多次）
//当我们第一次遍历字符串时，设当前遍历到的字符为 c，如果 c 不在哈希映射中，
//我们就将 c 与它的索引作为一个键值对加入哈希映射中，否则我们将 cc 在哈希映射中对应的值修改为−1。
