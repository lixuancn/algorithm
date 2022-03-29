package main

import "fmt"

//标题
//求最长不重复子串的长度
//
//题目描述
//str="abcabcdeabc"
//len=5 ("abcde")
func main() {
	str := "abcabcdeabc"
	i, j := 0, 0
	used := make(map[uint8]struct{})
	max := 0
	for ; j < len(str); j++ {
		//遇到重复
		if _, ok := used[str[i]]; ok {
			if max < j-i {
				fmt.Println(str[i : j+1])
				max = j - i
			}
			//移动i
			delete(used, str[i])
			i++
		} else {
			used[str[i]] = struct{}{}
		}
	}

	fmt.Println(max)
}
