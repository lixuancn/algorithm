package main

import "fmt"

//980. 不同路径 III

var ans int
var used [][]bool

func uniquePathsIII(grid [][]int) int {
	return uniquePathsIII_dfs(grid)
	//return uniquePathsIII_dp(grid) //动态规划比较难理解，而且还是得递归
}

func uniquePathsIII_dfs(grid [][]int) int {
	ans = 0
	used = make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		used[i] = make([]bool, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				backtracking(grid, i, j)
				return ans
			}
		}
	}
	return 0
}

func backtracking(grid [][]int, i, j int) {
	//关于判断这里，我是遍历，发现所有的used都被走过了，才算成功
	//题解很妙，有一个变量叫todo，初始化的时候就是只要不是-1，就todo++。在递归的时候，递归一次todo就减一，在这里判断的时候只需要判断todo是否为0
	if grid[i][j] == 2 {
		for m := 0; m < len(used); m++ {
			for n := 0; n < len(used[0]); n++ {
				if grid[m][n] == 0 && used[m][n] == false {
					return
				}
			}
		}
		ans++
		return
	}
	used[i][j] = true
	for _, pos := range [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		newI, newJ := i+pos[0], j+pos[1]
		if newI < 0 || newI >= len(grid) || newJ < 0 || newJ >= len(grid[0]) || used[newI][newJ] || grid[newI][newJ] == -1 {
			continue
		}
		backtracking(grid, newI, newJ)
	}
	used[i][j] = false
}

func main() {
	fmt.Println(uniquePathsIII([][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}}))
}
