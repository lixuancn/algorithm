package main

import "fmt"

//268. 丢失的数字

func missingNumber(nums []int) int {
	//return missingNumber_force(nums) //暴力法
	//return missingNumber_bit(nums) //位运算
	return missingNumber_math(nums) //数学法
	//return missingNumber_tag(nums) //打标记法
	//return missingNumber_swap(nums) //交换法
}

func missingNumber_force(nums []int) int {
	cache := make([]bool, len(nums)+1)
	for _, nums := range nums {
		cache[nums] = true
	}
	for i, ok := range cache {
		if !ok {
			return i
		}
	}
	return 0
}

//位运算，x^x=0, x^0=x的特性。先[0,n]进行异或，再循环nums，剩下的就是缺失的
func missingNumber_bit(nums []int) int {
	n := len(nums)
	result := 0
	for i := 0; i <= n; i++ {
		result ^= i
	}
	for i := 0; i < n; i++ {
		result ^= nums[i]
	}
	return result
}

//根据数学，高斯求和公式[0,n]的和是n*(n+1)/2。那么再减去nums的和，就是结果
func missingNumber_math(nums []int) int {
	n := len(nums)
	result := n * (n + 1) / 2
	for i := 0; i < n; i++ {
		result -= nums[i]
	}
	return result
}

//打标记法，如果存在一个数nums[i]，就把下标nums[i]-1给变成负数，也就是nums[nums[i]-1]变负。最后循环数组，如果一个数为正，那么他的下标这个数就是不存在的
func missingNumber_tag(nums []int) int {
	n := len(nums)
	//把0变成n，这样集合就成了[1,n+1]
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			nums[i] = n + 1
			break
		}
	}
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			nums[num-1] = -nums[num-1]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return 0
}

//把数字nums[i]放到nums[nums[i]-1]的位置上。由于交换最后会死循环，所以需要判断跳出
func missingNumber_swap(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] <= n && nums[i] > 0 && nums[i] != i+1 {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if i+1 != nums[i] {
			return i + 1
		}
	}
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))                   //2
	fmt.Println(missingNumber([]int{0, 1}))                      //2
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) //8
}
