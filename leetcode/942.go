package main

import "fmt"

//942. 增减字符串匹配

//贪心 如果s[0]=I，则perm[0]=0时，所有的perm[1]都可以满足
//如果s[0]=D，则perm[0]=n时，所有的perm[1]都可以满足
//然后这个问题就变为了长度为n-1的原问题
func diStringMatch(s string) []int {
	n := len(s)
	max, min := n, 0
	result := make([]int, n+1)
	k := 0
	for i := 0; i < n; i++ {
		if s[i] == 'I' {
			result[k] = min
			min++
		} else {
			result[k] = max
			max--
		}
		k++
	}
	result[k] = max
	return result
}

func main() {
	fmt.Println(diStringMatch("IDID")) //[0,4,1,3,2]
	fmt.Println(diStringMatch("III"))  //[0,1,2,3]
	fmt.Println(diStringMatch("DDI"))  //[3,2,0,1]
}
