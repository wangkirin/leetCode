package main

import "sort"

func main() {

}
func groupAnagrams(strs []string) [][]string {
	hash := make(map[string][]int)
	for i, str := range strs {
		bs := []byte(str)
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		sorted := string(bs)
		if _, ok := hash[sorted]; ok {
			hash[sorted] = append(hash[sorted], i)
		} else {
			hash[sorted] = []int{i}
		}
	}
	out := [][]string{}
	for _, v := range hash {
		tmp := []string{}
		for _, is := range v {
			tmp = append(tmp, strs[is])
		}
		out = append(out, tmp)
	}
	return out
}
