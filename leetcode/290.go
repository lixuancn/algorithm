package main

import (
	"fmt"
	"strings"
)

//290. 单词规律
//双映射
func wordPattern(pattern string, s string) bool {
	wordToChar := map[string]uint8{}
	charToWord := map[uint8]string{}
	wordList := strings.Split(s, " ")
	if len(pattern) != len(wordList) {
		return false
	}
	for i := range pattern {
		word := wordList[i]
		if _, ok := wordToChar[word]; ok && wordToChar[word] != pattern[i] {
			return false
		}
		if _, ok := charToWord[pattern[i]]; ok && charToWord[pattern[i]] != word {
			return false
		}
		wordToChar[word] = pattern[i]
		charToWord[pattern[i]] = word
	}
	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
	//fmt.Println(wordPattern("abba", "dog dog dog dog"))
	//fmt.Println(wordPattern("abba", "dog cat cat dog"))
}
