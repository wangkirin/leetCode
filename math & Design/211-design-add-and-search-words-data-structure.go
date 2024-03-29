package main

func main() {

}

//前缀树+DFS解法
type WordDictionary struct {
	isEnd bool
	next  [26]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{
		isEnd: false,
		next:  [26]*WordDictionary{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	node := this
	words := []byte(word)
	for _, b := range words {
		if node.next[b-'a'] == nil {
			node.next[b-'a'] = new(WordDictionary)
		}
		node = node.next[b-'a']
	}
	node.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	node := this
	words := []byte(word)
	var dfs func(i int, n *WordDictionary) bool
	dfs = func(i int, n *WordDictionary) bool {
		if i == len(words)-1 {
			return n.isEnd
		}
		c := words[i]
		if c != '.' {
			child := n.next[c-'a']
			if child != nil && dfs(i+1, child) {
				return true
			}
		} else {
			for _, dictionary := range n.next {
				if dictionary != nil && dfs(i+1, dictionary) {
					return true
				}
			}
		}
		return false
	}
	return dfs(0, node)
}

//暴力解法
//type WordDictionary struct {
//	words []string
//}
//
//func Constructor() WordDictionary {
//	return WordDictionary{words: nil}
//}
//
//func (this *WordDictionary) AddWord(word string) {
//	this.words = append(this.words, word)
//}
//
//func (this *WordDictionary) Search(word string) bool {
//	wordbytes := []byte(word)
//	count := len(word)
//	for _, s := range this.words {
//		flag := true
//		if count != len(s) {
//			continue
//		}
//		sbytes := []byte(s)
//		for i := 0; i < count; i++ {
//			if wordbytes[i] == '.' {
//				continue
//			}
//			if sbytes[i] != wordbytes[i] {
//				flag = false
//				break
//			}
//		}
//		if flag {
//			return true
//		}
//	}
//	return false
//}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
