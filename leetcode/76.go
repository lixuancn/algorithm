package main

import "fmt"

//76. 最小覆盖子串

//滑动窗口
func minWindow(s string, t string) string {
	oriRecord := map[byte]int{}
	curRecord := map[byte]int{}
	for i := range t {
		oriRecord[t[i]]++
	}
	var check = func() bool {
		for k, v := range oriRecord {
			if curRecord[k] < v {
				return false
			}
		}
		return true
	}
	ansL, ansR := -1, -1
	left, right := 0, 0
	minLen := 1 << 31
	for ; right < len(s); right++ {
		if _, ok := oriRecord[s[right]]; ok {
			curRecord[s[right]]++
		}
		for check() && left <= right {
			if right-left+1 < minLen {
				minLen = right - left + 1
				ansL, ansR = left, right
			}
			if _, ok := oriRecord[s[left]]; ok {
				curRecord[s[left]]--
			}
			left++
		}
	}
	if ansL == -1 || ansR == -1 {
		return ""
	}
	return s[ansL : ansR+1]
}

//暴力法
var record = map[byte]int{}

func minWindow_force(s string, t string) string {
	record = map[byte]int{}
	for i := range t {
		record[t[i]]++
	}
	ret, str := "", ""
	for i := 0; i <= len(s)-len(t); i++ {
		str = ""
		for j := i; j < len(s); j++ {
			str += string(s[j])
			if len(str) < len(t) || !isValid(s, i, j) {
				continue
			}
			if len(ret) == 0 || len(ret) > len(str) {
				ret = str
			}
		}
	}
	return ret
}

func isValid(s string, start, end int) bool {
	cache := make(map[byte]int)
	for i := start; i <= end; i++ {
		cache[s[i]]++
	}
	for char, count := range record {
		if v, ok := cache[char]; !ok || v < count {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(minWindow("a", "aa"))
	//fmt.Println(minWindow("abcd", "abc"))
	//fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
