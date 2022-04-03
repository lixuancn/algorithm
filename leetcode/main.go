package main

import "fmt"

var intToChinese = map[int]string{1: "一", 2: "二", 3: "三", 4: "四", 5: "五", 6: "六", 7: "七", 8: "八", 9: "九"}
var unitList = []string{"十", "百", "千", "万"}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[left] <= 9 {
			left = mid
		} else {
			right = mid - 1
		}
	}
	fmt.Println(left)
}
