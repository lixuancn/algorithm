package main

import "fmt"

var result [][]string
var row []string

func partition(s string) [][]string {
	result = make([][]string, 0)
	row = make([]string, 0)
	backtracking(s, 0)
	return result
}

func backtracking(s string, start int) {
	if start == len(s) {
		tmp := make([]string, len(row))
		copy(tmp, row)
		result = append(result, tmp)
	}
	for i := start; i < len(s); i++ {
		if isPalindrome(s, start, i) {
			row = append(row, s[start:i+1])
		} else {
			continue
		}
		backtracking(s, i+1)
		row = row[:len(row)-1]
	}
}

func isPalindrome(s string, left, right int) bool {
	if len(s) == 1 {
		return true
	}
	for i, j := left, right; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(partition("aab"))
}
