package main

import "fmt"

//滑动窗口
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	left, right := 0, n-1
	top, bottom := 0, n-1
	num := 1
	tar := n * n
	for num <= tar {
		for i := left; i <= right; i++ {
			fmt.Println(top, "-", i)
			matrix[top][i] = num
			num++
		}
		top++
		for i := top; i <= bottom; i++ {
			fmt.Println(top, "-", i)
			matrix[i][right] = num
			num++
		}
		right--
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}
	return matrix
}

func main() {
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(1))
	return
	//var nums []int
	//nums = []int{1, 2, 1}
	//fmt.Println(generateMatrix(nums))
	//nums = []int{0, 1, 2, 2}
	//fmt.Println(generateMatrix(nums))
	//nums = []int{1, 2, 3, 2, 2}
}
