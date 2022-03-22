package main

import "fmt"

//摩尔碰撞、摩尔投票法
//默认第一个数是一拨人，其他数是一拨人，循环是如果相同，就说明自己人，人数加一，如果不同，就自杀式干掉对方。
//最后count >= 1时说明是多数元素
func majorityElement(nums []int) int {
	count := 0
	ret := nums[0]
	for i, num := range nums {
		if num == ret {
			count++
		} else {
			count--
		}
		if count == 0 && i+1 < len(nums) {
			ret = nums[i+1]
		}
	}
	return ret
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3, 1, 0}))
}
