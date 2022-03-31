package main

import "fmt"

func firstUniqChar(s string) int {
	count := map[uint8]int{}
	for i := range s {
		count[s[i]]++
	}
	for i := range s {
		if count[s[i]] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("loveleetcode"))
}
