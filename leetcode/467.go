package main

import "fmt"

//467. 环绕字符串中唯一的子字符串

//dp[i]表示已p[i]结尾的最大子串长度。如，d结尾最大长度是bcd，那么长度有多长，子串就有多少个，bcd，cd，d
//遍历时，如果连续k++，不连续就置为1，然后赋值给dp[i]
//求得是唯一的子串长度，所以，abab中，a，b,a,b,ab,ab共六个只能成为a,b,ab共三个。故以某字母结尾的子串，我们取最长的即可
func findSubstringInWraproundString(p string) int {
	if len(p) < 1 {
		return 0
	}
	dp := [26]int{}
	k := 0
	for i := 0; i < len(p); i++ {
		//差是1或者-25，因为uint8，所以-25就是231
		if i > 0 && (p[i]-p[i-1] == 1 || p[i]-p[i-1] == 231) {
			k++
		} else {
			k = 1
		}
		index := p[i] - 'a'
		if k > dp[index] {
			dp[index] = k
		}
	}
	result := 0
	for _, v := range dp {
		result += v
	}
	return result
}

func main() {
	fmt.Println(findSubstringInWraproundString("a"))   //1
	fmt.Println(findSubstringInWraproundString("cac")) //2
	fmt.Println(findSubstringInWraproundString("zab")) //6
}
