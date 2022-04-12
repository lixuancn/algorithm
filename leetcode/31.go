package main

import "fmt"

//31. 下一个排列

//两次循环，看题解看不太懂，看代码加验算能看懂。
//比如13452的下一个是13524. 那就要新找到4和5，交换位置为13542，再把42翻转。
func nextPermutation(nums []int) {
	i, j := 0, 0
	//先找较小的数，例子中4
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		//再找较大的数，例子中5
		for j = len(nums) - 1; j >= 0; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	reverse(nums, i+1, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums) //1,3,2

	nums = []int{1, 3, 2}
	nextPermutation(nums)
	fmt.Println(nums) //2,1,3

	nums = []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums) //1,2,3

	nums = []int{1, 1, 5}
	nextPermutation(nums)
	fmt.Println(nums) //1,5,1
}
