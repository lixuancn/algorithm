package main

//73. 矩阵置零

//题解还有别的方法，可以优化掉空间的O(n)
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	zeroList := make([][2]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				zeroList = append(zeroList, [2]int{i, j})
			}
		}
	}
	usedRow := make(map[int]struct{}, 0)
	usedCol := make(map[int]struct{}, 0)
	for _, zeroPos := range zeroList {
		if _, ok := usedRow[zeroPos[0]]; !ok {
			for k := 0; k < n; k++ {
				matrix[zeroPos[0]][k] = 0
			}
		}
		if _, ok := usedCol[zeroPos[1]]; !ok {
			for k := 0; k < m; k++ {
				matrix[k][zeroPos[1]] = 0
			}
		}
		usedRow[zeroPos[0]] = struct{}{}
		usedCol[zeroPos[1]] = struct{}{}
	}
}
