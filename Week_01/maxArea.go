package main

import "fmt"

func maxArea(height []int) int {
	var max, area, left, right = 0, 0, 0, len(height)-1
	for ; left < right ;{
		if height[left] > height[right]{
			area = height[right] * (right - left)
			right--
		} else{
			area = height[left] * (right - left)
			left++
		}
		if area > max{
			max = area
		}
	}
	return max
}

func main(){
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}
