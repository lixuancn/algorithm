package main

import "fmt"

//240. 搜索二维矩阵 II

func searchMatrix(matrix [][]int, target int) bool {
	//return searchMatrix_mid(matrix, target) // 二分查找
	return searchMatrix_z(matrix, target) // z型查找
}

func searchMatrix_mid(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	for i := 0; i < n; i++ {
		if matrix[i][0] > target || target > matrix[i][m-1] {
			continue
		}
		for j := 0; j < len(matrix[i]); j++ {
			left, right := 0, len(matrix[i])-1
			for left <= right {
				mid := (left + right) >> 1
				if matrix[i][mid] == target {
					return true
				} else if matrix[i][mid] > target {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
		}
	}
	return false
}

//从右上角开始搜索，如果target大，则往下走一行（i++），如果target小，则往左走一格（j--）
//要么i++要么j--，所以最多n+m次，即O(n+m)
func searchMatrix_z(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	i, j := 0, m-1
	for i < n && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5))
	//fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20))
}
