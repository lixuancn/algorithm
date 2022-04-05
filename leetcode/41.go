package main

import "fmt"

//41. 缺失的第一个正数

func firstMissingPositive(nums []int) int {
	//return firstMissingPositive_hash(nums) //hash法
	return firstMissingPositive_swap(nums) //置换
}

//哈希法
//自建一个hash的话，空间复杂度就不是O(1)了。
//这里的思路是用数组模拟一个哈希。很妙，得看题解 https://leetcode-cn.com/problems/first-missing-positive/solution/que-shi-de-di-yi-ge-zheng-shu-by-leetcode-solution/
func firstMissingPositive_hash(nums []int) int {
	n := len(nums)
	//把非负整数变成正整数
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	//正整数x的下标定位x-1，为了标记x是存在的，就把x-1变成负数。
	//最后如果残缺的数在[1,n]的话，那么不是负数的数就是。
	//如果没有非负数，那么答案就是n+1
	//因为可能某个数被打过标记了（变成负的了），所以在遍历到的时候要取绝对值
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		//在[1,n]范围内的话
		if num <= n {
			nums[num-1] = -abs(nums[num-1])
		}
	}
	for i := 0; i < n; i++ {
		//有数没有被标记
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//置换法
//如果数组包含x，且x属于[1,n]，那么就把数字x放到下标x-1的位置
func firstMissingPositive_swap(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		//如果换过来的数还是属于[1,n]，那就继续换.
		//如果数字x和下标x-1是相同的话(数字x在他该在的位置上），就会死循环
		for nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{-1, 1, 3, 4, -2}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
}
