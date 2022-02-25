package main

import "fmt"

func removeElement_2(nums []int, val int) int {
	n := len(nums)
	left, right := 0, n-1
	for ; left<n; left++{
		if nums[left] == val{
			for ; right>left; right--{
				if nums[right] != val{
					break
				}
			}
			if left != right{
				nums[left], nums[right] = nums[right], nums[left]
			}
		}
		if right == left{
			break
		}
	}
	return right
}

func removeElement_force(nums []int, val int) int {
	n := len(nums)
	for i:=0; i<n; i++{
		if nums[i] == val{
			for j:=i+1; j<n; j++{
				nums[j-1] = nums[j]
			}
			i--
			n--
		}
	}
	return n
}

func main(){
	var nums []int
	nums = []int{2}
	fmt.Println(removeElement_force(nums, 3))
	fmt.Println(nums)
	nums = []int{3,2,2,3}
	fmt.Println(removeElement_force(nums, 3))
	fmt.Println(nums)
	nums = []int{0,1,2,2,3,0,4,2}
	fmt.Println(removeElement_force(nums, 2))
	fmt.Println(nums)
}