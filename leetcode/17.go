package main

import "fmt"

var ret []string
var retLine string
var numToLetter = map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}

func letterCombinations(digits string) []string {
	ret = make([]string, 0)
	if len(digits) != 0 {
		backracking(digits, 0)
	}
	return ret
}

func backracking(digits string, start int) {
	if len(digits) == len(retLine) {
		ret = append(ret, retLine)
		return
	}
	num := digits[start]
	if letters, ok := numToLetter[num]; ok {
		for _, letter := range letters {
			retLine += string(letter)
			backracking(digits, start+1)
			retLine = retLine[:len(retLine)-1]
		}
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
}
