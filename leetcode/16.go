package main

import (
	"fmt"
	"sort"
)

//16. 最接近的三数之和

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	return threeSumClosest_force(nums, target)
}

func threeSumClosest_force(nums []int, target int) int {
	result := 1 << 31
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			}
			if abs(sum-target) < abs(result-target) {
				result = sum
			}
			if sum > target {
				for k > 0 && nums[k] == nums[k-1] {
					k--
					continue
				}
				k--
			} else {
				for j < len(nums)-1 && nums[j] == nums[j+1] {
					j++
					continue
				}
				j++
			}
		}
	}
	return result
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1)) //2
	fmt.Println(threeSumClosest([]int{1, 1, 1, 0}, 100)) //3
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))      //0
	fmt.Println(threeSumClosest([]int{0, 2, 1, -3}, 1))  //0
}
