package main

import (
	"fmt"
	"sort"
)

//436. 寻找右区间
//注意，自己可以做自己的最右区间。测试用例就是这个干的

func findRightInterval(intervals [][]int) []int {
	//return findRightInterval_force(intervals) //暴力法
	return findRightInterval_mid(intervals)         //二分
	return findRightInterval_doubleIndex(intervals) //二分
}

//暴力法
func findRightInterval_force(intervals [][]int) []int {
	n := len(intervals)
	record := make([]int, n)
	for i := 0; i < n; i++ {
		record[i] = -1
		for j := 0; j < n; j++ {
			if intervals[i][1] <= intervals[j][0] {
				if record[i] == -1 || intervals[j][0] < intervals[record[i]][0] {
					record[i] = j
				}
			}
		}
	}
	return record
}

//二分 按起点排序
func findRightInterval_mid(intervals [][]int) []int {
	n := len(intervals)
	record := make([]int, n)
	//因为排序后，索引就变了，但结果要返回最初的索引
	for i := 0; i < n; i++ {
		intervals[i] = append(intervals[i], i)
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < n; i++ {
		pos := sort.Search(n, func(j int) bool {
			return intervals[i][1] <= intervals[j][0]
		})
		record[intervals[i][2]] = -1
		if pos < n {
			record[intervals[i][2]] = intervals[pos][2]
		}
	}
	return record
}

//双指针
//2个列表，一个是起始点，一个是结尾点
func findRightInterval_doubleIndex(intervals [][]int) []int {
	type pair struct {
		val   int
		index int
	}
	n := len(intervals)
	starts, ends := make([]pair, n), make([]pair, n)
	for i := 0; i < n; i++ {
		starts[i] = pair{val: intervals[i][0], index: i}
		ends[i] = pair{val: intervals[i][1], index: i}
	}
	sort.Slice(starts, func(i, j int) bool { return starts[i].val < starts[j].val })
	sort.Slice(ends, func(i, j int) bool { return ends[i].val < ends[j].val })
	record := make([]int, n)
	i := 0
	for j := range ends {
		//找开头比我的结尾还大的
		for i < n && starts[i].val < ends[j].val {
			i++
		}
		record[ends[j].index] = -1
		//找到了
		if i < n {
			record[ends[j].index] = starts[i].index
		}
	}
	return record
}

func main() {
	//fmt.Println(findRightInterval([][]int{{1, 2}}))                 //[-1]
	//fmt.Println(findRightInterval([][]int{{3, 4}, {2, 3}, {1, 2}})) //[-1,0,1]
	//fmt.Println(findRightInterval([][]int{{1, 4}, {2, 3}, {3, 4}})) //[-1,2,-1]
	//fmt.Println(findRightInterval([][]int{{1, 1}, {3, 4}}))         //[0,-1]
	fmt.Println(findRightInterval([][]int{{3, 4}, {2, 3}, {1, 2}})) //[-1,0,1]
}
