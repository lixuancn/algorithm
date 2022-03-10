package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	record := make(map[int32]int)
	for _, v := range magazine {
		record[v]++
	}
	for _, v := range ransomNote {
		if _, ok := record[v]; !ok {
			return false
		}
		record[v]--
	}
	for _, v := range record {
		if v < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}
