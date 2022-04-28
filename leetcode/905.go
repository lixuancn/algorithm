package main

import "fmt"

//905. 按奇偶排序数组

func sortArrayByParity(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; {
		if nums[j]&1 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		} else {
			j--
		}
	}
	return nums
}

//如果希望偶数顺序不要变
func sortArrayByParity_2(nums []int) []int {
	for i, j := 0, 0; j < len(nums); {
		if nums[j]&1 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		} else {
			j++
		}
	}
	return nums
}

//如果希望奇数顺序不要变
func sortArrayByParity_3(nums []int) []int {
	for i, j := len(nums)-1, len(nums)-1; j >= 0; {
		if nums[j]&1 == 1 {
			nums[i], nums[j] = nums[j], nums[i]
			i--
		}
		j--
	}
	return nums
}

//如果希望偶数和奇数顺序都不要变
//后来发现，之前写的不符合奇偶都不变得目标，但是也没想到怎么写更好。除了新开数组。
func sortArrayByParity_4(nums []int) []int {
	for i, j := 0, 0; j < len(nums) && i < len(nums); {
		//j先找到第一个偶数
		for ; j < len(nums); j++ {
			if nums[j]&1 == 0 {
				break
			}
		}
		//i找到第一个奇数
		for ; i <= j; i++ {
			if nums[i]&1 == 1 {
				break
			}
		}
		for i < len(nums) && j < len(nums) && nums[i]&1 == 1 && nums[j]&1 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}

//保持偶数顺序
func sortArrayByParity_practice(nums []int) []int {
	n := len(nums)
	ou := 0
	for i := 0; i < n; i++ {
		if nums[i]&1 == 0 {
			nums[i], nums[ou] = nums[ou], nums[i]
			ou++
		}
	}
	return nums
}

func main() {
	//fmt.Println(sortArrayByParity_practice([]int{3, 1, 2, 4}))
	//fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
	//fmt.Println(sortArrayByParity_2([]int{3, 1, 2, 4}))
	//fmt.Println(sortArrayByParity_3([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity_4([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity_4([]int{3, 1, 2}))
	fmt.Println(sortArrayByParity_4([]int{2, 4, 1, 3}))
}
