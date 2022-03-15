package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m && obstacleGrid[i][0] != 1; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n && obstacleGrid[0][j] != 1; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	//fmt.Println(uniquePathsWithObstacles([][]int{{0, 0}, {1, 1}, {0, 0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	//fmt.Println(uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}))
}
