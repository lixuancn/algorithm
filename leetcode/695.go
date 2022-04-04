package main

import (
	"container/list"
	"fmt"
)

//695. 岛屿的最大面积

func maxAreaOfIsland(grid [][]int) int {
	//return maxAreaOfIsland_dfs(grid)
	//return maxAreaOfIsland_dfs_stack(grid) //dfs把递归改成栈
	return maxAreaOfIsland_bfs(grid) //把栈改成队列，就是bfs了
}

var ans = 0
var area = 0

func maxAreaOfIsland_dfs(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	ans = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				area = 0
				maxAreaOfIsland_dfs_handler(grid, i, j)
			}
		}
	}
	return ans
}

func maxAreaOfIsland_dfs_handler(grid [][]int, i, j int) {
	grid[i][j] = 0
	area++
	if ans < area {
		ans = area
	}
	for _, pos := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		newI, newJ := i+pos[0], j+pos[1]
		if newI < 0 || newI >= len(grid) || newJ < 0 || newJ >= len(grid[0]) || grid[newI][newJ] != 1 {
			continue
		}
		maxAreaOfIsland_dfs_handler(grid, newI, newJ)
	}
}

func maxAreaOfIsland_dfs_stack(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	ans = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			area = 0
			stack := list.New()
			stack.PushBack([]int{i, j})
			for stack.Len() > 0 {
				cell := stack.Remove(stack.Back()).([]int)
				if cell[0] < 0 || cell[0] >= len(grid) || cell[1] < 0 || cell[1] >= len(grid[0]) || grid[cell[0]][cell[1]] != 1 {
					continue
				}
				area++
				grid[cell[0]][cell[1]] = 0
				for _, pos := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
					newI, newJ := cell[0]+pos[0], cell[1]+pos[1]
					stack.PushBack([]int{newI, newJ})
				}
			}
			if ans < area {
				ans = area
			}
		}
	}
	return ans
}

func maxAreaOfIsland_bfs(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	ans = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			area = 0
			queue := list.New()
			queue.PushBack([]int{i, j})
			for queue.Len() > 0 {
				cell := queue.Remove(queue.Front()).([]int)
				if cell[0] < 0 || cell[0] >= len(grid) || cell[1] < 0 || cell[1] >= len(grid[0]) || grid[cell[0]][cell[1]] != 1 {
					continue
				}
				area++
				grid[cell[0]][cell[1]] = 0
				for _, pos := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
					newI, newJ := cell[0]+pos[0], cell[1]+pos[1]
					queue.PushBack([]int{newI, newJ})
				}
			}
			if ans < area {
				ans = area
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maxAreaOfIsland([][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}))
}
