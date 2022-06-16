package main

import "fmt"

//944. 删列造序

func minDeletionSize(strs []string) int {
	result := 0
	for col := 0; col < len(strs[0]); col++ {
		for row := 1; row < len(strs); row++ {
			if strs[row][col] < strs[row-1][col] {
				result++
				break
			}
		}
	}
	return result
}

func main() {
	fmt.Println(minDeletionSize([]string{"cba", "daf", "ghi"})) //1
	fmt.Println(minDeletionSize([]string{"a", "b"}))            //0
	fmt.Println(minDeletionSize([]string{"zyx", "wvu", "tsr"})) //3
}
