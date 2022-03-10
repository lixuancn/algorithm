package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	record := make(map[int32]int)
	for _, c := range s {
		record[c]++
	}
	for _, c := range t {
		record[c]--
	}
	for _, v := range record {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}
