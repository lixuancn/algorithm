package main

//48. 旋转图像

func rotate(matrix [][]int) {
	//rotate_rotate_extra(matrix) //旋转，用额外数组
	//rotate_rotate(matrix) //旋转，不用额外数组
	rotate_overturn(matrix) //翻转，先水平翻转，在对角线翻转
}

//我们会发现，一个点matrix[row][col]旋转后在matrix[col][n-1-row]
func rotate_rotate_extra(matrix [][]int) {
	n := len(matrix)
	newMatrix := make([][]int, n)
	for row := 0; row < n; row++ {
		newMatrix[row] = make([]int, n)
	}
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			newMatrix[col][n-1-row] = matrix[row][col]
		}
	}
	copy(matrix, newMatrix)
}

//我们会发现，一个点matrix[row][col]旋转后在matrix[col][n-1-row]，如果直接复制，那么这个点的值就会被覆盖掉。
//而我们发现，一个正方形，旋转4次就会回到原点，那么就一次循环完成四次赋值
//matrix[row][col] -> matrix[col][n-1-row]
//matrix[col][n-1-row] -> matrix[n-1-row][n-1-col]
//matrix[n-1-row][n-1-col] -> matrix[n-1-col][n-1-(n-1-row)]=matrix[n-1-col][row]
//matrix[n-1-col][row] -> matrix[row][n-1-(n-1-col)]=matrix[row][col]
//所以我们用一个临时变量来解决覆盖值的问题。在go里就更简单了，连tmp都不需要了
//matrix[col][n-1-row],matrix[n-1-row][n-1-col],matrix[n-1-col][row],matrix[row][col] = matrix[row][col],matrix[col][n-1-row],matrix[n-1-row][n-1-col],matrix[n-1-col][row]
//至于循环次数，一共有n*n个点，我们一次换4个，那么偶数个点时，需要交换n*n/4次，奇数时，中心点是不需要动的，所以需要(n*n-1)/4次
//比如4*4的方格，我们需要把[0,1]行[0,1]列进行旋转四次(n/2和n/2)，比如5*5的方格，需要把[0,1]行[0,2]列进行旋转四次(n-1/2，n+1/2)
func rotate_rotate(matrix [][]int) {
	n := len(matrix)
	for row := 0; row < n/2; row++ {
		for col := 0; col < (n+1)/2; col++ {
			matrix[col][n-1-row], matrix[n-1-row][n-1-col], matrix[n-1-col][row], matrix[row][col] = matrix[row][col], matrix[col][n-1-row], matrix[n-1-row][n-1-col], matrix[n-1-col][row]
		}
	}
}

//翻转，先水平翻转，在主对角线翻转
//比如[[1,2,3],[4,5,6],[7,8,9]]，水平对折后是[[7,8,9],[4,5,6],[1,2,3]]，主对角线对折（左上到右下的对角线）后[[7,4,1],[8,5,2],[9,6,3]]，即答案
//推导：matrix[row][col]水平对折后是matrix[n-1-row][col]，主对角线对折后是matrix[col][n-1-row]。
//这就和之前的公式matrix[row][col]旋转后在matrix[col][n-1-row]一致
func rotate_overturn(matrix [][]int) {
	n := len(matrix)
	//水平对折-注意循环终点，否则又换回去了
	for row := 0; row < n/2; row++ {
		matrix[row], matrix[n-1-row] = matrix[n-1-row], matrix[row]
	}
	//主对角线对折，注意col的终点，否则又换回去了
	for row := 0; row < n; row++ {
		for col := 0; col < row; col++ {
			matrix[row][col], matrix[col][row] = matrix[col][row], matrix[row][col]
		}
	}
}
}
