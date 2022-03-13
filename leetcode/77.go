package main

import "fmt"

var res [][]int
var resLine []int

func combine(n int, k int) [][]int {
	res = make([][]int, 0)
	resLine = make([]int, 0)
	backtracking(n, k, 1)
	return res
}

func backtracking(n, k, start int) {
	if len(resLine) == k {
		tmp := make([]int, k)
		copy(tmp, resLine)
		res = append(res, tmp)
		return
	}
	//这里可以剪纸，比如我还需要3个数，但只剩2个了，就不用继续循环了
	//for i := start; i <= n; i++ {
	for i := start; i <= n-(k-len(resLine))+1; i++ {
		resLine = append(resLine, i)
		backtracking(n, k, i+1)
		resLine = resLine[:len(resLine)-1]
	}
}

func main() {
	fmt.Println(combine(1, 1))
}
