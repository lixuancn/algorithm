package main

import "fmt"

//切片拼接 或 自行创建切片填充
//字符串是不可变类型，每次循环都要新申请空间创建新字符串变量，所以不如上面的切片类型

func main() {
	r := reverseLeftWords("abcdefg", 2)
	fmt.Println(r)
}

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[0:n]
}

func reverseLeftWords_loop(s string, n int) string {
	l := len(s)
	var r []byte
	//for i:=n; i<l; i++ {
	//	r = append(r, s[i])
	//}
	//for i:=0; i<n; i++ {
	//	r = append(r, s[i])
	//}
	//可简化为：
	for i:=n; i<l+n; i++{
		r = append(r, s[i % l])
	}
	return string(r)
}
