package main

import (
	"container/list"
	"fmt"
)

//200. 岛屿数量

//遍历地图中的1，把遍历过的1都改成0
func numIslands(grid [][]byte) int {
	//return numIslands_dfs(grid)
	return numIslands_bfs(grid)
}

func numIslands_dfs(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				ans++
				numIslands_dfs_handler(grid, i, j)
			}
		}
	}

	return ans
}

func numIslands_dfs_handler(grid [][]byte, i, j int) {
	grid[i][j] = '0'
	for _, pos := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		newI, newJ := i+pos[0], j+pos[1]
		if newI < 0 || newI >= len(grid) || newJ < 0 || newJ >= len(grid[0]) || grid[newI][newJ] != '1' {
			continue
		}
		numIslands_dfs_handler(grid, newI, newJ)
	}
}

func numIslands_bfs(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				ans++
				queue := list.New()
				queue.PushBack([]int{i, j})
				for queue.Len() > 0 {
					cell := queue.Remove(queue.Front()).([]int)
					grid[cell[0]][cell[1]] = '0'
					for _, pos := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
						newI, newJ := cell[0]+pos[0], cell[1]+pos[1]
						if newI < 0 || newI >= len(grid) || newJ < 0 || newJ >= len(grid[0]) || grid[newI][newJ] != '1' {
							continue
						}
						queue.PushBack([]int{newI, newJ})
					}
				}
			}
		}
	}
	return ans
}

func main() {
	//fmt.Println(numIslands([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}))
	fmt.Println(numIslands([][]byte{{'1', '0', '1', '1', '0', '1', '1'}}))
}
