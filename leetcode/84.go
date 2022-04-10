package main

import "fmt"

//84. 柱状图中最大的矩形

func largestRectangleArea(heights []int) int {
	//return largestRectangleArea_force_weight(heights) //暴力法，枚举宽
	//return largestRectangleArea_force_height(heights) //暴力法，枚举高
	return largestRectangleArea_stack(heights) //单调栈
}

//暴力法，枚举宽
func largestRectangleArea_force_weight(heights []int) int {
	result, area := 0, 0
	for i := 0; i < len(heights); i++ {
		minHeight := heights[i]
		for j := i; j < len(heights); j++ {
			if minHeight > heights[j] {
				minHeight = heights[j]
			}
			area = (j - i + 1) * minHeight
			if result < area {
				result = area
			}
		}
	}
	return result
}

//暴力法，枚举高。 以一根柱子为高，双指针向左向右寻找边际，如果下一个比自己的低的，那么上一个就是边界了
func largestRectangleArea_force_height(heights []int) int {
	result, area := 0, 0
	for i, height := range heights {
		left, right := i, i
		for left-1 >= 0 && heights[left-1] >= height {
			left--
		}
		for right+1 < len(heights) && heights[right+1] >= height {
			right++
		}
		area = (right - left + 1) * height
		if result < area {
			result = area
		}
	}
	return result
}

//单调栈 详情看题解
func largestRectangleArea_stack(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	monoStack := []int{}
	for i := 0; i < n; i++ {
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			left[i] = -1
		} else {
			left[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	monoStack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			right[i] = n
		} else {
			right[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	result := 0
	for i := 0; i < n; i++ {
		area := (right[i] - left[i] - 1) * heights[i]
		if result < area {
			result = area
		}
	}
	return result
}

func main() {
	fmt.Println(largestRectangleArea([]int{6, 7, 5, 2, 4, 5, 9, 3})) //16
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))       //10
	fmt.Println(largestRectangleArea([]int{2, 4}))                   //4
}
