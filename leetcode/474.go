package main

import "fmt"

func findMaxForm(strs []string, m int, n int) int {
	return findMaxForm_dp(strs, m, n) //动态规划
	max := 0
	findMaxForm_recursion(strs, m, n, 0, 0, &max) //回溯
	return max
}

func findMaxForm_dp(strs []string, m int, n int) int {
	if len(strs) == 0 {
		return 0
	}
	//dp[i][j] 表示有i个0，j个1的最大子集数
	//如果不放，就是dp[i][j], 如果放，就是dp[i-zeroNum][j-oneNum]+1，取最大值
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	//循环物品
	for _, str := range strs {
		zeroNum, oneNum := getOneOrZeroCount(str)
		//循环背包容量，倒序
		for i := m; i >= zeroNum; i-- {
			for j := n; j >= oneNum; j-- {
				if dp[i][j] < dp[i-zeroNum][j-oneNum]+1 {
					dp[i][j] = dp[i-zeroNum][j-oneNum] + 1
				}
			}
		}
	}
	return dp[m][n]
}

//回溯 会超时
func findMaxForm_recursion(strs []string, m, n, start, count int, max *int) {
	if m == 0 || n == 0 || start == len(strs) {
		if *max < count {
			*max = count
		}
	}
	for i := start; i < len(strs); i++ {
		countM, countN := getOneOrZeroCount(strs[i])
		if countM > m || countN > n {
			continue
		}
		if m-countM < 0 || n-countN < 0 {
			continue
		}
		count++
		findMaxForm_recursion(strs, m-countM, n-countN, i+1, count, max)
		count--
	}
}

func getOneOrZeroCount(str string) (int, int) {
	countM, countN := 0, 0
	for _, c := range str {
		if c == '1' {
			countN++
		} else if c == '0' {
			countM++
		}
	}
	return countM, countN
}

func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
}
