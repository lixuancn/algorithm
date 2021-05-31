package main
import (
	"fmt"
)

func main() {
	r := isPowerOfFour(16)
	fmt.Println(r)
}

//位运算，二进制表示后，有且只有一个1，且1在奇数位(从1开始数）
//n是32位置有符号整数，那么新建一个mask，他的奇数位都是1，偶数位都是0。n是4次幂的时候，n & mask一定为n
func isPowerOfFour(n int) bool {
	return n > 0 && n&(n-1)==0 && n&0b1010101010101010101010101010101==n
}

//位运算，找规律发现，二进制表示后，有且只有一个1，且1在奇数位(从1开始数）。
func isPowerOfFour_bit(n int) bool {
	if n <= 0 || n & (n-1) != 0{
		return false
	}
	for n >= 1{
		if n & 1 == 1{
			return true
		}
		n = n >> 2
	}
	return false
}

//直接算-除法
func isPowerOfFour_violence_division(n int) bool {
	if n <= 0{
		return false
	}
	cur := float32(n)
	for ; cur>1; cur=cur/4{
		if cur == 1{
			return true
		}
	}
	return cur==1
}

//直接算-乘法
func isPowerOfFour_violence_multiplication(n int) bool {
	if n <= 0{
		return false
	}
	if n == 1{
		return true
	}
	cur := 1
	for cur < n{
		if cur *= 4; cur == n{
			return true
		}
	}
	return false
}
