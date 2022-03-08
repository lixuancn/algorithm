package main

import "fmt"

//滑动窗口
func totalFruit(fruits []int) int {
	record := make(map[int]int)
	left, right, size := 0, 0, 0
	for ; right < len(fruits); right++ {
		record[fruits[right]]++
		for len(record) >= 3 {
			record[fruits[left]]--
			if record[fruits[left]] <= 0 {
				delete(record, fruits[left])
			}
			left++
		}
		if size < right-left+1 {
			size = right - left + 1
		}
	}
	return size
}

func main() {
	var nums []int
	nums = []int{1, 2, 1}
	fmt.Println(totalFruit(nums))
	nums = []int{0, 1, 2, 2}
	fmt.Println(totalFruit(nums))
	nums = []int{1, 2, 3, 2, 2}
	fmt.Println(totalFruit(nums))
	nums = []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	fmt.Println(totalFruit(nums))
}
