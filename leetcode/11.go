package main

import "fmt"

//11. 盛最多水的容器

func maxArea(height []int) int {
	result, area, left, right := 0, 0, 0, len(height)-1
	for left < right {
		if height[left] < height[right] {
			area = height[left] * (right - left)
			left++
		} else {
			area = height[right] * (right - left)
			right--
		}
		if result < area {
			result = area
		}
	}
	return result
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) //49
	fmt.Println(maxArea([]int{1, 1}))                      //1
}
