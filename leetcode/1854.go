package main

import "fmt"

//给你一个二维整数数组 logs ，其中每个 logs[i] = [birthi, deathi] 表示第 i 个人的出生和死亡年份。
//年份 x 的 人口 定义为这一年期间活着的人的数目。第 i 个人被计入年份 x 的人口需要满足：x 在闭区间 [birthi, deathi - 1] 内。注意，人不应当计入他们死亡当年的人口中。
//返回 人口最多 且 最早 的年份。
func main() {
	var logs = [][]int{{1950,1961},{1960,1971},{1970,1981}}

	logs = [][]int{{2033,2034},{2039,2047},{1998,2042},{2047,2048},{2025,2029},{2005,2044},{1990,1992},{1952,1956},{1984,2014}}

	r := maximumPopulation(logs)
	fmt.Println(r)
}

//差分数组法
//每个个体对人口数量的影响，只在出生和死亡这两个年份中。
//最后，在给定人口初值与每一年人口变化量的基础上，我们可以将对应的变化量求和得到每一年的人口数量，进而得到人口最多的年份。
//而将变化量转换为对应数量的过程正是求解「前缀和」的方法，因此「差分」也是「前缀和」的逆运算。
func maximumPopulation(logs [][]int) int {
	var yearList [101]int
	for _, personal := range logs{
		yearList[personal[0]-1950]++
		yearList[personal[1]-1950]--
	}
	var max, maxYear, curr int
	for k, v := range yearList{
		curr += v
		if curr > max{
			max = curr
			maxYear = k
		}
	}
	return maxYear + 1950
}

func maximumPopulation_loop(logs [][]int) int {
	var yearList [100]int
	for _, personal := range logs{
		for i:=personal[0]-1950; i<=personal[1]-1-1950; i++{
			yearList[i]++
		}
	}
	max := 0
	maxYear := 0
	for k, v := range yearList {
		if max < v{
			max = v
			maxYear = k + 1950
		}
	}
	return maxYear
}
