package main
import (
	"fmt"
	"strconv"
	"math/bits"
)

func main() {
	r := hammingDistance_byself(1, 4)
	fmt.Println(r)
}

//Brian Kernighan算法
func hammingDistance_BrianKernighan(x, y int) (ans int) {
	for s := x ^ y; s > 0; s &= s - 1 {
		ans++
	}
	return
}

//发明轮子发
func hammingDistance_byself(x, y int) int {
	ans := 0
	for s := x ^ y ; s > 0; s>>=1{
		//得到s的最低位
		ans += s & 1
	}
	return ans
}

//内置函数法
func hammingDistance_inside(x, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

//暴力法
func hammingDistance_violence(x int, y int) int {
	xBin := int32ToBin(x)
	yBin := int32ToBin(y)
	n := len(xBin)
	if len(yBin) > n{
		n = len(yBin)
	}

	xBin = fmt.Sprintf("%0*s", n, xBin)
	yBin = fmt.Sprintf("%0*s", n, yBin)
	distance := 0
	for i:=0; i<n; i++{
		if xBin[i] != yBin[i]{
			distance++
		}
	}
	fmt.Println(xBin)
	fmt.Println(yBin)
	return distance
}

func int32ToBin(num int)string{
	s := ""
	if num == 0{
		return "0"
	}
	for num>0{
		s = strconv.Itoa(num%2) + s
		num /= 2
	}
	return s
}
