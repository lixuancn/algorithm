package main

import "fmt"

//6. Z 字形变换

func convert(s string, numRows int) string {
	n := len(s)
	if numRows == 1 || n <= numRows {
		return s
	}
	//一个周期是向下numRows，再向上numRows-2。所以一个周期是numRows+numRows-2
	t := 2*numRows - 2
	//我们有n/t个周期，每个周期有numRows-2+1列。所以有
	c := (n + t - 1) / t * (numRows - 1)
	mat := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		mat[i] = make([]byte, c)
	}
	x, y := 0, 0
	for i, ch := range s {
		mat[x][y] = byte(ch)
		if i%t < numRows-1 {
			x++ //向下
		} else {
			x-- //向右上
			y++
		}
	}
	result := make([]byte, n)
	k := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] > 0 {
				result[k] = mat[i][j]
				k++
			}
		}
	}
	return string(result)
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3)) //PAHNAPLSIIGYIR
	fmt.Println(convert("PAYPALISHIRING", 4)) //PINALSIGYAHRPI
}
