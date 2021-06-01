package main
import (
	"fmt"
)

//给你一个二进制字符串 s 。如果字符串中由 1 组成的 最长 连续子字符串 严格长于 由 0 组成的 最长 连续子字符串，返回 true ；否则，返回 false 。
//例如，s = "110100010" 中，由 1 组成的最长连续子字符串的长度是 2 ，由 0 组成的最长连续子字符串的长度是 3 。
//注意，如果字符串中不存在 0 ，此时认为由 0 组成的最长连续子字符串的长度是 0 。字符串中不存在 1 的情况也适用此规则。

func main() {
	r := checkZeroOnes("1101")
	fmt.Println(r)
}

func checkZeroOnes(s string) bool {
	maxOne, maxZero := 0, 0
	var last byte
	lastLen := 0
	for i,n:=0,len(s); i<=n; i++{
		if i == n{
			if last == '1' && maxOne < lastLen{
				maxOne = lastLen
			}else if last == '0' && maxZero < lastLen{
				maxZero = lastLen
			}
			break
		}
		if last != s[i]{
			if last == '1' && maxOne < lastLen{
				maxOne = lastLen
			}else if last == '0' && maxZero < lastLen{
				maxZero = lastLen
			}
			last = s[i]
			lastLen = 1
		}else{
			lastLen++
		}
	}
	return maxOne > maxZero
}
