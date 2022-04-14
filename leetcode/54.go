package main

import "fmt"

//54. 螺旋矩阵

func spiralOrder(matrix [][]int) []int {
	//return spiralOrder_simulation(matrix) //模拟
	return spiralOrder_simulation_layer(matrix) //按层模拟
}

//模拟
func spiralOrder_simulation(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	result := make([]int, m*n)
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} //顺时针旋转
	directionIndex := 0
	row, col := 0, 0
	for i := 0; i < m*n; i++ {
		result[i] = matrix[row][col]
		visited[row][col] = true
		//遇到边就换方向
		nextRow, nextCol := row+directions[directionIndex][0], col+directions[directionIndex][1]
		if nextRow < 0 || nextRow == m || nextCol < 0 || nextCol == n || visited[nextRow][nextCol] {
			directionIndex = (directionIndex + 1) % 4
		}
		row += directions[directionIndex][0]
		col += directions[directionIndex][1]
	}
	return result
}

//模拟 - 按层模拟，先[top,left]到[top,right]，再[top+1,right][bottom,right]，再[bottom,right-1]到[bottom,left+1]，再[bottom,left]到[top+1,left]
//然后bottom-1,top+1,left+1,right-1，开始第二圈的循环
func spiralOrder_simulation_layer(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	result := make([]int, m*n)
	index := 0
	left, right, top, bottom := 0, n-1, 0, m-1
	fmt.Println(left, right, top, bottom)
	for left <= right && top <= bottom {
		for col := left; col <= right; col++ {
			result[index] = matrix[top][col]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			result[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			for col := right - 1; col >= left+1; col-- {
				result[index] = matrix[bottom][col]
				index++
			}
			for row := bottom; row >= top+1; row-- {
				result[index] = matrix[row][left]
				index++
			}
		}
		top++
		bottom--
		left++
		right--
	}
	return result
}

func main() {
	//fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(spiralOrder([][]int{{3}, {2}}))
}
