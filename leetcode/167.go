package main

//167. 两数之和 II - 输入有序数组

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}
