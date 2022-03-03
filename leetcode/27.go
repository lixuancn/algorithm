package main

import "fmt"

//双指针指针
func removeElement_doubleIndex(nums []int, val int) int {
	n := len(nums)
	left, right := 0, n-1
	for ; left <= right; left++ {
		if nums[left] == val {
			nums[left] = nums[right]
			left--
			right--
		}
	}
	return left
}

//快慢指针
func removeElement_fastAndSlowIndex(nums []int, val int) int {
	n := len(nums)
	slow, fast := 0, 0
	for ; fast < n; fast++ {
		if val != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func removeElement_force(nums []int, val int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == val {
			for j := i + 1; j < n; j++ {
				nums[j-1] = nums[j]
			}
			i--
			n--
		}
	}
	return n
}

func main() {
	var nums []int
	nums = []int{2}
	fmt.Println(removeElement_fastAndSlowIndex(nums, 3))
	fmt.Println(nums)
	nums = []int{3, 2, 2, 3}
	fmt.Println(removeElement_fastAndSlowIndex(nums, 3))
	fmt.Println(nums)
	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement_fastAndSlowIndex(nums, 2))
	fmt.Println(nums)
	nums = []int{0, 1, 2, 2, 3, 0, 4, 1}
	fmt.Println(removeElement_fastAndSlowIndex(nums, 2))
	fmt.Println(nums)
}
