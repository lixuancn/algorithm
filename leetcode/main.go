package main

import "fmt"

//输入[2, 4, 5, 1, 6 ,8 ,13]，单次循环，不是用额外空间，将偶数移动到数组尾部，且不改变奇数顺序
func main() {
	nums := []int{1, 2, 4, 5, 1, 6, 8, 13}
	move(nums)
	//[5 1 13 2 4 6 8]
	fmt.Println(nums)
}

func move(nums []int) {
	left, right := 0, 1
	for right < len(nums) && left < len(nums) {
		//如果是偶数
		if nums[left]&1 == 0 {
			nums[left], nums[right] = nums[right], nums[left]
			right++
			if nums[left]&1 != 0 {
				left++
			}
		} else {
			left++
			right++
		}
	}
}
