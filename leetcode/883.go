package main

import "fmt"

//883. 三维形体投影面积

func projectionArea(grid [][]int) int {
	n := len(grid)
	result := 0
	//第一个平面，就是数组的元素个数之和
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				result++
			}
		}
	}
	//第二个平面，取每个一维数组中的最大值，比如grid[i][0](grid[0][0],grid[1][0],grid[2][0])的最大值，grid[i][1]的最大值
	for j := 0; j < n; j++ {
		m2 := 0
		for i := 0; i < n; i++ {
			if m2 < grid[i][j] {
				m2 = grid[i][j]
			}
		}
		result += m2
	}
	//第三个平面，取同纬度的最大值，比如grid[0][j](grid[0][0],grid[0][1],grid[0][2])的最大值，grid[1][j]的最大值
	for i := 0; i < n; i++ {
		m3 := 0
		for j := 0; j < n; j++ {
			if m3 < grid[i][j] {
				m3 = grid[i][j]
			}
		}
		result += m3
	}
	return result

	//代码简化
	n := len(grid)
	result := 0
	for i := 0; i < n; i++ {
		m2, m3 := 0, 0
		for j := 0; j < n; j++ {
			//第一个平面，就是数组的元素个数之和
			if grid[i][j] > 0 {
				result++
			}
			//第二个平面，取每个一维数组中的最大值，比如grid[i][0](grid[0][0],grid[1][0],grid[2][0])的最大值，grid[i][1]的最大值
			if m2 < grid[j][i] {
				m2 = grid[j][i]
			}
			//第三个平面，取同纬度的最大值，比如grid[0][j](grid[0][0],grid[0][1],grid[0][2])的最大值，grid[1][j]的最大值
			if m3 < grid[i][j] {
				m3 = grid[i][j]
			}
		}
		result += m2
		result += m3
	}
	return result
}

func main() {
	fmt.Println(projectionArea([][]int{{1, 2}, {3, 4}})) //17
	fmt.Println(projectionArea([][]int{{2}}))            //5
	fmt.Println(projectionArea([][]int{{1, 0}, {0, 2}})) //8
}
