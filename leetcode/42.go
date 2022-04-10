package main

import (
	"fmt"
)

//42. 接雨水

func trap(height []int) int {
	//return trap_doublePointer(height) //双指针
	return trap_dp(height) //动态规划
}

//遍历每一列。选中一个柱子，然后向两边找大于他的柱子。那么这一列能装的水是min(left,right)*1-height[i]
func trap_doublePointer(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	result := 0
	for i := 1; i < len(height)-1; i++ {
		leftHeight, rightHeight := height[i], height[i]
		for left := i - 1; left >= 0; left-- {
			if height[left] > leftHeight {
				leftHeight = height[left]
			}
		}
		for right := i + 1; right < len(height); right++ {
			if height[right] > rightHeight {
				rightHeight = height[right]
			}
		}
		result += min(leftHeight, rightHeight)*1 - height[i]
	}
	return result
}

//我们发现，在双指针的解法中，每次都需要向左右去寻找最高的柱子，如果记录下的话，就可以减少循环。
//leftDp[i]表示[0,i]柱子的最高的高度，leftDp[i] = max(leftDp[i-1], height[i])
//rightDp[i]表示[i,n]柱子的最高的高度，rightDp[i] = max(rightDp[i-1], height[i])
func trap_dp(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	leftDp := make([]int, len(height))
	leftDp[0] = height[0]
	rightDp := make([]int, len(height))
	rightDp[len(height)-1] = height[len(height)-1]
	for i := 1; i < len(height); i++ {
		leftDp[i] = max(leftDp[i-1], height[i])
	}
	for i := len(height) - 2; i >= 0; i-- {
		rightDp[i] = max(rightDp[i+1], height[i])
	}
	fmt.Println(leftDp)
	fmt.Println(rightDp)
	result := 0
	for i := 1; i < len(height)-1; i++ {
		area := min(leftDp[i-1], rightDp[i+1]) - height[i]
		if area > 0 {
			result += area
		}
	}
	return result
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) //6
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))                   //9
	fmt.Println(trap([]int{4, 2, 3}))                            // 1
}
