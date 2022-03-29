package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	i, j := -1, 0
	used := make(map[uint8]int)
	max := 0
	for ; j < len(s); j++ {
		//遇到重复，i直接跳到重复字段的下标
		if _, ok := used[s[j]]; ok {
			if i < used[s[j]] {
				i = used[s[j]]
			}
		}
		used[s[j]] = j
		if max < j-i {
			max = j - i
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abba"))
	//fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
