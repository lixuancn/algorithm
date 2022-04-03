package main

import "fmt"

//74. 搜索二维矩阵

//两次二分
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	left, right := 0, len(matrix)-1
	for left <= right {
		mid := (left + right) >> 1
		if matrix[mid][0] == target || matrix[mid][len(matrix[mid])-1] == target {
			return true
		}
		if matrix[mid][0] < target && target < matrix[mid][len(matrix[mid])-1] {
			i, j := 0, len(matrix[mid])-1
			for i <= j {
				m := (i + j) >> 1
				if target == matrix[mid][m] {
					return true
				}
				if target < matrix[mid][m] {
					j = m - 1
				} else {
					i = m + 1
				}
			}
		}
		if matrix[mid][0] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

//题解里有1次二分，是把矩阵连成一个有序数组。
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		mid := (left + right) >> 1
		x := matrix[mid/n][mid%n]
		if x < target {
			left = mid + 1
		} else if x > target {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
}
