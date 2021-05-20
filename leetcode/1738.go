package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//给你一个二维矩阵 matrix 和一个整数 k ，矩阵大小为 m x n 由非负整数组成。
//矩阵中坐标 (a, b) 的 值 可由对所有满足 0 <= i <= a < m 且 0 <= j <= b < n 的元素 matrix[i][j]（下标从 0 开始计数）执行异或运算得到。
//请你找出 matrix 的所有坐标中第 k 大的值（k 的值从 1 开始计数）。

//前缀和 s[i] = matrix[0]^matrix[1]^...^matrix[i]
//那么，二维前缀和 s[i][j] = matrix[0][0]^matrix[0][1]^...^matrix[0][j]
//              ^matrix[1][0]^matrix[1][1]^...^matrix[1][j]
//              ^matrix[i][0]^matrix[i][1]^...^matrix[i][j]
//对于(a,b)，而(a,b)和(i,j)概念是一样的。  就是求s[i][j]

func main() {
	param := [][]int{{5,2},{1,6}}
	param = [][]int{{8,10,5,8,5,7,6,0,1,4,10,6,4,3,6,8,7,9,4,2}}
	r := kthLargestValue(param, 2)
	fmt.Println(r)
}

//使用排序
func kthLargestValue(matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])
	var s = make([][]int, m)
	for i:=0; i<m; i++{
		colValue := 0
		for j:=0; j<n; j++{
			colValue ^= matrix[i][j]
			s[i] = append(s[i], colValue)
		}
	}
	var ret = make([]int, 0)
	for i:=0; i<m; i++{
		for j:=0; j<n; j++{
			if i > 0{
				s[i][j] ^= s[i-1][j]
			}
			ret = append(ret, s[i][j])
		}
	}
	//比排序更低复杂度
	return quickSelect(ret, len(ret)-k)
	sort.Ints(ret)
	return ret[len(ret)-k]
}

//比排序更低复杂度
func quickSelect(a []int, k int) int {
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	for l, r := 0, len(a)-1; l < r; {
		v := a[l]
		i, j := l, r+1
		for {
			for i++; i < r && a[i] < v; i++ {
			}
			for j--; j > l && a[j] > v; j-- {
			}
			if i >= j {
				break
			}
			a[i], a[j] = a[j], a[i]
		}
		a[l], a[j] = a[j], v
		if j == k {
			break
		} else if j < k {
			l = j + 1
		} else {
			r = j - 1
		}
	}
	return a[k]
}

func kthLargestValue_loop(matrix [][]int, k int) int {
	var ret []int
	for a := range matrix{
		for b := range matrix[a]{
			r := 0
			for i:=0; i<=a; i++{
				for j:=0; j<=b; j++{
					r ^= matrix[i][j]
				}
			}
			ret = append(ret, r)
		}
	}
	sort.Ints(ret)
	return ret[len(ret)-k]
}
