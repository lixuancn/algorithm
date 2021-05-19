package main

import (
	"fmt"
)

//给你一个整数数组 arr 。
//现需要从数组中取三个下标 i、j 和 k ，其中 (0 <= i < j <= k < arr.length) 。
//a 和 b 定义如下：
//a = arr[i] ^ arr[i + 1] ^ ... ^ arr[j - 1]
//b = arr[j] ^ arr[j + 1] ^ ... ^ arr[k]
//注意：^ 表示 按位异或 操作。
//请返回能够令 a == b 成立的三元组 (i, j , k) 的数目。

//输入：arr = [2,3,1,6,7]
//输出：4
//解释：满足题意的三元组分别是 (0,1,2), (0,2,2), (2,3,4) 以及 (2,4,4)

//S为异或前缀和，Si = arr[0] ^ ... ^ arr[i] = Si-1 ^ arr[i]
//那么Si ^ Sj = arr[i+1] ^ ... ^ arr[j]
//所以数组[i, j]的异或和等于arr[i] ^ ... ^ arr[j] = Si-1 ^ Sj
//那么a = Si-1 ^ Sj-1
//那么b = Sj-1 ^ Sk
//a=b时，则Si-1 ^ Sj-1 = Sj-1 ^ Sk 即 Si-1 = Sk
//由于i从0开始，i-1有下标不存在问题，所以转为Si = Sk+1
func main() {
	param := []int{2,3,1,6,7}
	//param = []int{1,1,1,1,1}
	r := countTriplets_loop_2(param)
	fmt.Println(r)
}

//三重循环
func countTriplets_loop_3(arr []int) int {
	length := len(arr)
	var s = make([]int, length+1)
	for k := range arr{
		s[k+1] = s[k] ^ arr[k]
	}
	count := 0
	for i:=0; i<length; i++{
		for j:=i+1; j<length; j++{
			for k:=j; k<length; k++ {
				if s[i] == s[k+1]{
					count++
				}
			}
		}
	}
	return count
}

//二重循环
//当Si=Sk+1时，即在范围[i+1, k]内任意数为j都是成立的，则j有k-i个
func countTriplets_loop_2(arr []int) int {
	length := len(arr)
	var s = make([]int, length+1)
	for k := range arr{
		s[k+1] = s[k] ^ arr[k]
	}
	count := 0
	for i:=0; i<length; i++{
		for k:=i+1; k<length; k++ {
			if s[i] == s[k+1]{
				count += k - i
			}
		}
	}
	return count
}

//一重循环
func countTriplets_loop_1(arr []int) int {
	length := len(arr)
	var s = make([]int, length+1)
	for k := range arr{
		s[k+1] = s[k] ^ arr[k]
	}
	cnt := map[int]int{}
	total := map[int]int{}
	count := 0
	for k:=0; k<length; k++{
		if m, ok := cnt[s[k+1]]; ok{
			count += m * k - total[s[k+1]]
		}
		cnt[s[k]]++
		total[s[k]] += k
	}
	return count
}

func countTriplets(arr []int) int {
	return countTriplets_loop_1(arr)
}
