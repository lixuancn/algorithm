package main

import (
	"fmt"
)

//668. 乘法表中第k小的数

//二分
//第i行的所有数，都是i的倍数。
//假设有x，如果x>i*n 那么这一行比x小的就有n个。
//如果x<i*n，且x在这一行，那么这一行比x小的等的就x/i
//如果x<i*n，且x不在这一行，那么这一行比x小的就是x/i
func findKthNumber(m int, n int, k int) int {
	if k == 1 {
		return 1
	} else if k == m*n {
		return m * n
	}
	if k < 1 || k > m*n {
		return -1
	}
	//获取比x小的数有多少个，即x是第几小
	var getSmallerThanX = func(x int) int {
		result := 0
		for i := 1; i <= n; i++ {
			if i*m < x {
				result += m
			} else {
				result += x / i
			}
		}
		return result
	}
	left, right := 1, m*n
	for left < right {
		mid := (left + right) >> 1
		count := getSmallerThanX(mid)
		if count < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}

func main() {
	fmt.Println(findKthNumber(3, 3, 5)) //3
	fmt.Println(findKthNumber(2, 3, 6)) //6
}
