package main

import "fmt"

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
func sortArrayByParity_4(nums []int) []int {
	for i, j := 0, 0; j < len(nums) && i < len(nums); {
		fmt.Println(nums)
		//j先找到第一个偶数
		for ; j < len(nums); j++ {
			if nums[j]&1 == 0 {
				break
			}
		}
		//i找到第一个奇数
		for ; i < len(nums); i++ {
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

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity_2([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity_3([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity_4([]int{3, 1, 2, 4}))
}
