package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	cache := make(map[uint8]int)
	max := 0
	left, right := 0, 0
	for right = 0; right < n; right++ {
		c := s[right]
		if _, ok := cache[c]; ok {
			left = maxInt(left, cache[c]+1)
		}
		cache[c] = right
		max = maxInt(max, right-left+1)
	}
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lengthOfLongestSubstring_practice("abba"))
	//fmt.Println(lengthOfLongestSubstring_practice("abcabcbb"))
	//fmt.Println(lengthOfLongestSubstring_practice("bbbbb"))
	//fmt.Println(lengthOfLongestSubstring_practice("pwwkew"))
}

func lengthOfLongestSubstring_practice(s string) int {
	left, right := -1, 0
	l := 0
	used := map[uint8]int{}
	for ; right < len(s); right++ {
		if _, ok := used[s[right]]; ok {
			left = maxInt(left, used[s[right]])
		}
		used[s[right]] = right
		fmt.Println(right, left)
		l = maxInt(l, right-left)
	}
	return l
}
