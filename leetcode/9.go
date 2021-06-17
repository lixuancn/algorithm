package main

import "fmt"

func main() {
	r := isPalindrome(121)
	fmt.Println(r)
}

//只翻转一半
func isPalindrome(x int) bool {
	//负数、个位是0的 不可能是回文
	if x < 0 || (x!=0 && x%10==0){
		return false
	}
	y := 0
	for x > y{
		y = y*10 + x%10
		x /= 10
	}
	//偶数位数可以直接比，奇数位数需要去掉中间的数后再比
	return x == y || x == y/10
}

//整个数字翻转，可能面临超出MaxInt的溢出问题
func isPalindrome_all(x int) bool {
	if x < 0{
		return false
	}
	i, j := x, 0
	for i > 0{
		j = j*10 + i%10
		i = i / 10
	}
	return j == x
}
