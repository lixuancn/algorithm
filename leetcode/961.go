package main

import "fmt"

//961. 在长度 2N 的数组中找出重复 N 次的元素

//方法一，模拟
func repeatedNTimes(nums []int) int {
	cache := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := cache[num]; ok {
			return num
		}
		cache[num] = struct{}{}
	}
	return -1
}

//方法二，数学
//假设，两个x之间至少隔了2个数字，那么总长度就是n + (n-1)*2 = 3n-2 //n个重复数，有n-1个间隙
//n > 2时， 3n-2 > 2n，不符合题意
//n = 2时， 2n=4，所以最多是间隔2，
//遍历间隔，嵌套遍历nums，看这两个数是否相等即可，即nums[i] == nums[gap+i]

//方法三，随机
//[0,2n-1]随机两个数，如果不等，但是num相等，则说明找到了x。
//概率是四分之一，即四次就能找到

func main() {
	fmt.Println(repeatedNTimes([]int{1, 2, 3, 3}))
	fmt.Println(repeatedNTimes([]int{2, 1, 2, 5, 3, 2}))
	fmt.Println(repeatedNTimes([]int{5, 1, 5, 2, 5, 3, 5, 4}))
}
