package main

func main() {

}

//前缀树

type Trie struct {
	isEnd bool
	next  [26]*Trie
}

func Constructor() Trie {
	return Trie{
		isEnd: false,
		next:  [26]*Trie{},
	}
}

func (this *Trie) Insert(word string) {
	words := []byte(word)
	node := this
	for _, b := range words {
		if node.next[b-'a'] == nil {
			node.next[b-'a'] = new(Trie)
		}
		node = node.next[b-'a']
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {

	words := []byte(word)
	node := this
	for _, b := range words {
		node = node.next[b-'a']
		if node == nil {
			return false
		}
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	prefixs := []byte(prefix)
	node := this
	for _, b := range prefixs {
		node = node.next[b-'a']
		if node == nil {
			return false
		}
	}
	return true
}

//暴力解法
//type Trie struct {
//	wordsMap map[string]int
//}
//
//func Constructor() Trie {
//	return Trie{map[string]int{}}
//}
//
//func (this *Trie) Insert(word string) {
//	this.wordsMap[word] = 1
//}
//
//func (this *Trie) Search(word string) bool {
//	if _, ok := this.wordsMap[word]; ok {
//		return true
//	}
//	return false
//}
//
//func (this *Trie) StartsWith(prefix string) bool {
//	for k, _ := range this.wordsMap {
//		if strings.HasPrefix(k, prefix) {
//			return true
//		}
//	}
//	return false
//}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
