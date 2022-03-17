package main

import (
	"fmt"
	"strconv"
	"strings"
)

var result [][]string
var row []string

func restoreIpAddresses(s string) []string {
	result = make([][]string, 0)
	row = make([]string, 0)
	backtracking(s, 0)
	ret := make([]string, len(result))
	for i := 0; i < len(result); i++ {
		ret[i] = strings.Join(result[i], ".")
	}
	return ret
}

func backtracking(s string, start int) {
	if start == len(s) {
		return
	}
	if len(row) == 3 {
		if !isValidItem(s, start, len(s)) {
			return
		}
		tmp := make([]string, 4)
		tmp[0], tmp[1], tmp[2], tmp[3] = row[0], row[1], row[2], s[start:]
		result = append(result, tmp)
		return
	}
	for i := start; i < len(s); i++ {
		if !isValidItem(s, start, i+1) {
			break
		}
		row = append(row, s[start:i+1])
		backtracking(s, i+1)
		row = row[:len(row)-1]
	}
}

func isValidItem(s string, left, right int) bool {
	if s[left] == '0' && right-left == 1 {
		return true
	} else if s[left] == '0' && right-left > 1 {
		return false
	}
	num, err := strconv.Atoi(s[left:right])
	if err != nil || num > 255 {
		return false
	}
	return true
}

func main() {
	fmt.Println(restoreIpAddresses("101023"))
}
