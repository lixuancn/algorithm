package main

import "fmt"

var res = [][]int{}
var resLine = []int{}
var sum = 0

func combinationSum3(k int, n int) [][]int {
	res = make([][]int, 0)
	backtracking(k, n, 1)
	return res
}

func backtracking(k, n, start int) {
	if sum == n && len(resLine) == k {
		tmp := make([]int, k)
		copy(tmp, resLine)
		res = append(res, tmp)
		return
	} else if sum >= n || len(resLine) >= k {
		return
	}
	for i := start; i <= 9-(k-len(resLine))+1; i++ {
		sum += i
		resLine = append(resLine, i)
		backtracking(k, n, i+1)
		sum -= i
		resLine = resLine[:len(resLine)-1]
		if sum+i > n {
			break
		}
	}
}

func main() {
	fmt.Println(combinationSum3(3, 7))
}
