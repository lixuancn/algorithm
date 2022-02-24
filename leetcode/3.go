package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0{
		return 0
	}
	cache := make(map[uint8]int)
	max := 0
	left, right := 0, 0
	for right=0; right<n; right++{
		c := s[right]
		if _, ok := cache[c]; ok{
			left = maxInt(left, cache[c]+1)
		}
		cache[c] = right
		max = maxInt(max, right-left+1)
	}
	return max
}

func maxInt(a, b int)int{
	if a > b {
		return a
	}
	return b
}

func main(){
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}