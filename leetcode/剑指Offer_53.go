package main

import "fmt"

//剑指 Offer 53 - II. 0～n-1中缺失的数字

func missingNumber(nums []int) int {
	//return missingNumber_map(nums)  //方法一：map
	//return missingNumber_sum(nums)  //方法二：数学法加和
	//return missingNumber_xor(nums)  //方法三：异或
	return missingNumber_tag(nums)  //方法四：标记法
	return missingNumber_swap(nums) //方法五：交换法
}

//方法一：map
func missingNumber_map(nums []int) int {
	n := len(nums)
	cache := make([]int, n+1)
	for i := 0; i < n; i++ {
		cache[nums[i]] = 1
	}
	for i := 0; i < n+1; i++ {
		if cache[i] == 0 {
			return i
		}
	}
	return n
}

//方法二：数学法加和
func missingNumber_sum(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for _, num := range nums {
		sum -= num
	}
	return sum
}

//方法三：异或
func missingNumber_xor(nums []int) int {
	n := len(nums)
	xor := 0
	for i := 0; i < n+1; i++ {
		xor ^= i
	}
	for _, num := range nums {
		xor ^= num
	}
	return xor
}

//方法四：标记法
//把数组模拟成一个map，由此可以不用开辟额外空间
//数字nums[i]如果存在，就把nums[nums[i]-1]变成负数。最后遍历nums，如果为正，则这个下标+1这个数不存在
func missingNumber_tag(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			nums[num-1] = -nums[num-1]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return 0
}

//方法五：交换法
//和方法四类似，也是把数组当成map，只不过不是打标记，而是做交换，把nums[i]放在下标nums[i]-1的下标上。
//最后循环nums，看下标i+1是不是数字nums[i]
func missingNumber_swap(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i] != i+1 {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return 0
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	fmt.Println(missingNumber([]int{1, 3, 2}))
	fmt.Println(missingNumber([]int{0, 1, 2}))
	fmt.Println(missingNumber([]int{0, 1, 3}))
	fmt.Println(missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 7, 9}))
}
